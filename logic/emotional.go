package logic

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"graduation/dao/mysql"
	"graduation/domain"
	"graduation/models"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"go.uber.org/zap"
)

func Emotional(ctx context.Context, p *models.EmotionalText) (*domain.EmotionalReply, error) {
	// 请求情感分析接口
	replychan := make(chan *domain.EmotionalReply, 1)
	errapichan := make(chan error, 1)

	go func() {
		defer close(replychan)
		defer close(errapichan)
		reply, err := getEmotionalInfofromali(p)
		if err != nil {
			errapichan <- err
			return
		}
		replychan <- reply
	} ()

	// 控制请求阿里的接口超时
	var reply *domain.EmotionalReply
	select {
	case reply = <-replychan:
		break
	case err := <-errapichan:
		zap.L().Error("sendemotionInfo failed", zap.Error(err))
		return nil, err
	case <-ctx.Done():
		zap.L().Error("context timeout", zap.Error(ctx.Err()))
		return nil, ctx.Err()
	}

	go func ()  {
		// 保存到数据库
		dbctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
		defer cancel()
		err := mysql.SaveEmotionInfo(dbctx, p, reply)
		if err != nil {
			zap.L().Error("mysql.Emotional failed", zap.Error(err))
		}
	} ()

	return reply, nil
}

func GetEmotionalInfo(ctx context.Context, userID int64) (*models.EmotionalInfo, error) {
	return mysql.GetEmotionalInfo(ctx, userID)
}

// 请求情感分析接口函数
func sendemotionInfo(ctx context.Context, p *models.EmotionalText) (*domain.EmotionalReply, error) {
	// 定义请求数据
	reqData := domain.EmotionalReq{
		Text: p.Text,
	}
	jsonData, err := json.Marshal(reqData)
	if err != nil {
		zap.L().Error("json.Marshal failed", zap.Error(err))
		return nil, err
	}

	// 构造请求体和请求头
	req, err := http.NewRequestWithContext(ctx, "POST", domain.EmotionalUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		zap.L().Error("http.NewRequest failed", zap.Error(err))
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	// 发送请求
	client := &http.Client{Timeout: domain.HttpTimeout}
	resp, err := client.Do(req)
	if err != nil {
		zap.L().Error("client.Do failed", zap.Error(err))
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		zap.L().Error("ioutil.ReadAll failed", zap.Error(err))
		return nil, err
	}

	var reply domain.EmotionalReply
	err = json.Unmarshal(body, &reply)
	if err != nil {
		zap.L().Error("json.Unmarshal failed", zap.Error(err))
		return nil, err
	}
	return &reply, nil
}

func getEmotionalInfofromali(p *models.EmotionalText) (*domain.EmotionalReply, error) {
	// 请求阿里情感分析接口
	client, err := sdk.NewClientWithAccessKey("cn-hangzhou", os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"))
    if err != nil {
        panic(err)
    }

	request := requests.NewCommonRequest()
	request.Domain = "alinlp.cn-hangzhou.aliyuncs.com"
    request.Version = "2020-06-29"
	request.ApiName = "GetSaChGeneral"
    request.QueryParams["ServiceCode"] = "alinlp"
    request.QueryParams["Text"] = p.Text
    // request.QueryParams["TokenizerId"] = "GENERAL_CHN"
    response, err := client.ProcessCommonRequest(request)
    if err != nil {
        panic(err)
    }

	// 解析请求结果
	jsonStr := string(response.GetHttpContentBytes())
    var resp domain.Response
	err = json.Unmarshal([]byte(jsonStr), &resp)
	if err != nil {
		return nil, err
	}

	// 解析 Data 部分
	var data domain.DataContent
	err = json.Unmarshal([]byte(resp.Data), &data)
	if err != nil {
		fmt.Println("解析 Data JSON 失败:", err)
		return nil, err
	}
	// 打印情绪分析结果
	fmt.Printf("\n情感分析结果：%s \n正面情绪的概率为：%.4f \n负面情绪的概率为：%.4f \n中性情绪的概率为：%.4f \n",
		data.Result.Sentiment, data.Result.PositiveProb, data.Result.NegativeProb, data.Result.NeutralProb)

	var result domain.EmotionalReply
	result.Text = p.Text
	result.SentimentType = domain.GetSentimentType()[data.Result.Sentiment]
	result.Sentiment = domain.GetSentimentMsg()[data.Result.Sentiment]

	return &result, nil
}
