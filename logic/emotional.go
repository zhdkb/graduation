package logic

import (
	"bytes"
	"context"
	"encoding/json"
	"graduation/dao/mysql"
	"graduation/domain"
	"graduation/models"
	"io/ioutil"
	"net/http"
	"time"

	"go.uber.org/zap"
)

func Emotional(ctx context.Context, p *models.EmotionalText) (*domain.EmotionalReply, error) {
	// 请求情感分析接口
	reply, err := sendemotionInfo(ctx, p)
	if err != nil {
		zap.L().Error("sendemotionInfo failed", zap.Error(err))
		return nil, err
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
