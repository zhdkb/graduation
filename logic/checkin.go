package logic

import (
	"context"
	"graduation/dao/mysql"
	"graduation/domain"

	"go.uber.org/zap"
)

func CheckIn(ctx context.Context, check *domain.CheckIn) (string, error) {
	streak_days, err := mysql.GetCheckInCount(ctx, check.UserID, check.CheckInTime)
	if err != nil {
		zap.L().Error("GetCheckInCount failed", zap.Error(err))
		return "", err
	}

	err = mysql.StoreCheckin(ctx, check, streak_days + 1)
	if err != nil {
		zap.L().Error("StoreCheckin failed", zap.Error(err))
		return "", err
	}

	return "打卡成功", nil
}
