package mysql

import (
	"context"
	"errors"
	"graduation/domain"
	"graduation/models"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

// GetCheckInCount 查询用户的打卡记录
func GetCheckInCount(ctx context.Context, userID int64, checkInTime time.Time) (int, error) {
	db := db.Model(&models.CheckInRecord{})
	// 去数据库中查询前一天的连续打卡天数
	var checkInRecord models.CheckInRecord
	lastDay := checkInTime.AddDate(0, 0, -1).Format("2006-01-02")
	db.Where("check_in_time < ? and check_in_time >= ?", checkInTime, lastDay).Where("user_id = ?", userID)
	err := db.Find(&checkInRecord).Error
	if err != nil {
		// if errors.Is(err, gorm.ErrRecordNotFound) {
		// 	// 如果没有找到记录，说明今天是第一次打卡，连续打卡天数为0
		// 	return 0, nil
		// }
		zap.L().Error("GetCheckInCount failed", zap.Error(err))
		return 0, err
	}

	if checkInRecord.CheckInTime.Format("2006-01-02") == checkInTime.Format("2006-01-02") {
		// 如果今天已经打卡了，返回不要重复打卡
		return 0, errors.New("今天已经打卡了")
	}

	return checkInRecord.StreakDays, nil
}

// StoreCheckin 存储打卡记录和更新打卡统计信息
func StoreCheckin(ctx context.Context, checkIn *domain.CheckIn, streak_days int) error {
	// 开始一个事务
	tx := db.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 调用 StoreCheckinRecord 函数插入打卡记录
	if err := StoreCheckinRecord(ctx, checkIn.UserID, checkIn.CheckInTime, streak_days); err != nil {
		tx.Rollback() // 如果插入失败，回滚事务
		return err
	}

	// 调用 StoreCheckinCounts 函数更新打卡统计信息
	if err := StoreCheckinCounts(ctx, checkIn.UserID, streak_days); err != nil {
		tx.Rollback() // 如果更新失败，回滚事务
		return err
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}

// StoreCheckinRecord 存储打卡记录
func StoreCheckinRecord(ctx context.Context, userID int64, checkInTime time.Time, streakDays int) error {
	db := db.Model(&models.CheckInRecord{})
	checkIn := &models.CheckInRecord{
		UserID:      userID,
		CheckInTime: checkInTime,
		StreakDays:  streakDays,
		CreateTime: time.Now(),
	}
	return db.WithContext(ctx).Create(checkIn).Error
}

// StoreCheckinCounts 更新打卡统计信息
func StoreCheckinCounts(ctx context.Context, userID int64, streakDays int) error {
	db := db.Model(&models.CheckInCount{})
	// 首先查询用户是否存在
	var checkInCount models.CheckInCount
	err := db.WithContext(ctx).Where("user_id = ?", userID).First(&checkInCount).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// 如果不存在，则插入新的记录
		checkInCount := &models.CheckInCount{
			UserID:        userID,
			CheckInCount:  1,
			MaxStreakDays: streakDays,
			CreateTime:    time.Now(),
			UpDateTime:    time.Now(),
		}
		if errCreate := db.WithContext(ctx).Create(checkInCount).Error; errCreate != nil {
			zap.L().Error("StoreCheckinCounts failed", zap.Error(errCreate))
			return errCreate
		}
	} else if err != nil {
		zap.L().Error("StoreCheckinCounts failed", zap.Error(err))
		return err
	} else {
		// 如果存在，则更新打卡次数和连续打卡天数
		checkInCount.CheckInCount++
		if streakDays > checkInCount.MaxStreakDays {
			checkInCount.MaxStreakDays = streakDays
		}
		if errUpdate := db.WithContext(ctx).Where("user_id = ?", userID).Updates(map[string]interface{} {
			"check_in_count":  checkInCount.CheckInCount,
			"max_streak_days": checkInCount.MaxStreakDays,
			"update_time":    time.Now(),
		}).Error; errUpdate != nil {
			zap.L().Error("StoreCheckinCounts failed", zap.Error(errUpdate))
			return errUpdate
		}
	}

	return nil
}
