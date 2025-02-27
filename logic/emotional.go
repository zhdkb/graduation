package logic

import (
	"context"
	"graduation/dao/mysql"
	"graduation/models"
)

func Emotional(ctx context.Context, p *models.EmotionalText) (*models.EmotionalResponse, error) {
	return mysql.Emotional(ctx, p)
}
