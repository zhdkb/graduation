package mysql

import (
	"context"
	"graduation/domain"
	"graduation/models"
	"time"

	"go.uber.org/zap"
)

func Emotional(ctx context.Context, p *models.EmotionalText) (*models.EmotionalResponse, error) {
	sqlStr := "insert into emotional(user_id, text, msg, data) values(?, ?, ?, ?)"
	err := db.WithContext(ctx).Exec(sqlStr, p.UserID, p.Text, "good", "12312").Error
	if err != nil {
		return nil, err
	}
	
	return &models.EmotionalResponse{
		Msg: "good",
	}, nil
}

func GetEmotionalInfo(ctx context.Context, userID int64) (*models.EmotionalInfo, error) {
	sqlStr := "count(if(emotional_type = 1, 1, null)) as good_num, count(if(emotional_type = 2, 1, null)) as bad_num, count(if(emotional_type = 0, 1, null)) as neutral_num"
	var info models.EmotionalInfo
	db := db.Model(&models.EmotionalInfo{})
	err := db.WithContext(ctx).Select(sqlStr).Where("user_id = ?", userID).Scan(&info).Error
	if err != nil {
		return nil, err
	}

	return &info, nil
}

func SaveEmotionInfo(ctx context.Context, p *models.EmotionalText, reply *domain.EmotionalReply) error {
	emotionalInfo := domain.EmotionalInfo{
		UserID:        p.UserID,
		EmotionalText: p.Text,
		EmotionalType: reply.SentimentType,
		EmotionalMsg:  reply.Sentiment,
		CreateTime:    time.Now(),
		ModifyTime:    time.Now(),
	}

	err := db.WithContext(ctx).Create(&emotionalInfo).Error
	if err != nil {
		zap.L().Error("SaveEmotionInfo failed", zap.Error(err))
	}
	return err
}
