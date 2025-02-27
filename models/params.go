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
	UserID int64  `json:"user_id" binding:"required"`
    Text string `json:"text" binding:"required"`
}

type EmotionalResponse struct {
	Msg  string `json:"msg"`
    Data string `json:"data"`
}

type EmotionalInfo struct {
	GoodNum int64  `json:"good_num"`
	BadNum 	int64  `json:"bad_num"`
	NeutralNum int64  `json:"neutral_num"`
}
