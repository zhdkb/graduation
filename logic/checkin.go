package logic

import (
	"context"
	"graduation/dao/mysql"
	"graduation/dao/redis"
	"graduation/domain"
	"time"

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

	go func () {
		newctx, cancel := context.WithTimeout(context.Background(), 2 * time.Second)
		defer cancel()
		err := redis.SetCheckInCount(newctx, check.CheckInTime)
		if err != nil {
			zap.L().Error("今日打卡总数存入redis失败", zap.Error(err))
			return
		}
	} ()

	return "打卡成功", nil
}
