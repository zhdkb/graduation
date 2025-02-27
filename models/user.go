package models

type User struct {
	UserID		int64	`gorm:"column:user_id"`
	Username	string	`gorm:"column:username"`
	Password	string	`gorm:"column:password"`
	Token		string
}

type UserID struct {
	UserID int64	`json:"user_id"`
}
