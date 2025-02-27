package mysql

import (
	"context"
	"graduation/models"
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
	sqlStr := "select count(if(msg = 'good', 1, null)) as good_num, count(if(msg = 'bad', 1, null)) as bad_num, count(if(msg = 'neutral', 1, null)) as neutral_num from emotional where user_id = ?"
	var info models.EmotionalInfo
	err := db.WithContext(ctx).Raw(sqlStr, userID).Scan(&info).Error
	if err != nil {
		return nil, err
	}

	return &info, nil
}
