package domain

import "time"

type EmotionalInfo struct {
	Id            int64  `json:"id" gorm:"column:id"`
	UserID        int64  `json:"user_id" gorm:"column:user_id"`
	EmotionalText string `json:"emotional_text" gorm:"column:emotional_text"`
	EmotionalType int    `json:"emotional_type" gorm:"column:emotional_type"`
	EmotionalMsg  string `json:"emotional_msg" gorm:"column:emotional_msg"`
	CreateTime    time.Time `json:"create_time" gorm:"column:create_time"`
	ModifyTime	  time.Time `json:"modify_time" gorm:"column:modify_time"`
}

func (EmotionalInfo) TableName() string {
	return "emotional_infos"
}
