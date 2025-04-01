package domain

// 外层结构体
type Response struct {
	RequestId string `json:"RequestId"`
	Data      string `json:"Data"` // Data 先作为字符串解析
}

// 内层结构体
type DataContent struct {
	Result struct {
		PositiveProb float64 `json:"positive_prob"`
		Sentiment    string  `json:"sentiment"`
		NeutralProb  float64 `json:"neutral_prob"`
		NegativeProb float64 `json:"negative_prob"`
	} `json:"result"`
	Success  bool   `json:"success"`
	TracerId string `json:"tracerId"`
}

const (
	PositiveSentimentType = 1 // 正向情感
	NegativeSentimentType = 2 // 负向情感
	NeutralSentimentType  = 0 // 中性情感

	PositiveMsg      = "这句话表达了正向情感"
	NegativeMsg      = "这句话表达了负向情感"
	NeutralMsg       = "这句话表达了中性情感"

	PositiveSentiment = "正面"
	NegativeSentiment = "负面"
	NeutralSentiment  = "中性"
)

func GetSentimentType() map[string]int {
	return map[string]int {
		PositiveSentiment: PositiveSentimentType,
		NegativeSentiment: NegativeSentimentType,
		NeutralSentiment:  NeutralSentimentType,
	}
}

func GetSentimentMsg() map[string]string {
	return map[string]string {
		PositiveSentiment: PositiveMsg,
		NegativeSentiment: NegativeMsg,
		NeutralSentiment:  NeutralMsg,
	}
}
