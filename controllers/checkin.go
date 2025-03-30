package controllers

import (
	"graduation/domain"
	"graduation/logic"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// CheckinHandler 实现打卡逻辑
func CheckinHandler(c *gin.Context) {
	// 获取用户ID
	userIDStr := c.Param("user_id")
	userid, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "userID获取失败",
		})
		return
	}

	check := &domain.CheckIn{
		UserID:      int64(userid),
		CheckInTime: time.Now(),
	}

	// 调用逻辑层处理打卡逻辑
	result, err := logic.CheckIn(c.Request.Context(), check)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": result,
		"msg":  "success",
	})

}
