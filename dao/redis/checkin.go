package redis

import (
	"context"
	"time"
)

// SetCheckInCount 存储今日的打卡总数到 Redis
func SetCheckInCount(ctx context.Context, checkInTime time.Time) error {
	key := checkInTime.Format("2006-01-02")
	_, err := rdb.Incr(ctx, key).Result()

	return err
}
