package logic

import (
	"context"
	"graduation/dao/mysql"
	"graduation/models"
)

func Emotional(ctx context.Context, p *models.EmotionalText) (*models.EmotionalResponse, error) {
	return mysql.Emotional(ctx, p)
}

func GetEmotionalInfo(ctx context.Context, userID int64) (*models.EmotionalInfo, error) {
	return mysql.GetEmotionalInfo(ctx, userID)
}
