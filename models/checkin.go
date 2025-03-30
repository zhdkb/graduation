package models

import "time"

type CheckInRecord struct {
	ID          int64     `json:"id" gorm:"column:id"`
	UserID      int64     `json:"user_id" gorm:"column:user_id"`
	CheckInTime time.Time `json:"check_in_time" gorm:"column:check_in_time"`
	StreakDays  int       `json:"streak_days" gorm:"column:streak_days"`
	CreateTime  time.Time `json:"create_time" gorm:"column:create_time"`
}

type CheckInCount struct {
	UserID        int64     `json:"user_id" gorm:"column:user_id"`
	CheckInCount  int       `json:"check_in_count" gorm:"column:check_in_count"`
	MaxStreakDays int       `json:"max_streak_days" gorm:"column:max_streak_days"`
	CreateTime    time.Time `json:"create_time" gorm:"column:create_time"`
	UpDateTime    time.Time `json:"update_time" gorm:"column:update_time"`
}

func (CheckInRecord) TableName() string {
	return "check_in_records"
}

func (CheckInCount) TableName() string {
	return "user_check_in_counts"
}
