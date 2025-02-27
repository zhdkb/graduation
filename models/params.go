package models

// 定义请求的参数结构体

const (
	OrderTime = "time"
	OrderScore = "score"
)

// ParamSignUp 注册请求参数
type ParamSignUp struct {
	Username	string	`json:"username" binding:"required"`
	Password	string	`json:"password" binding:"required"`
	RePassword	string	`json:"re_password" binding:"required,eqfield=Password"`
}

// ParamLogin 登录请求参数
type ParamLogin struct {
	Username	string	`json:"username" binding:"required"`
	Password	string	`json:"password" binding:"required"`
}

type EmotionalText struct {
    Text string `json:"text" binding:"required"`
}

type EmotionalResponse struct {
    Data string `json:"data"`
}
