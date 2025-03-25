package domain

import "time"

const (
	EmotionalUrl = "http://127.0.0.1:5000/predict"
	HttpTimeout  = 5 * time.Second
)

type EmotionalReq struct {
	Text string `json:"text"`
}

type EmotionalReply struct {
	Sentiment     string `json:"sentiment"`
	SentimentType int    `json:"sentiment_type"`
	Text          string `json:"text"`
}
