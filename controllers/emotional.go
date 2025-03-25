package controllers

import (
	"graduation/logic"
	"graduation/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func EmotionalHandler(c *gin.Context) {
	p := new(models.EmotionalText)
	if err := c.ShouldBindJSON(&p); err != nil {
		//请求参数有误，返回响应
		zap.L().Error("Emotional with invalid param", zap.Error(err))

		c.JSON(http.StatusOK, gin.H{
			"msg": err,
		})
		return
	}

	result, err := logic.Emotional(c.Request.Context(), p)

	if err != nil {
		zap.L().Error("logic.Emotional failed", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": result,
		"msg": "success",
	})

}

func EmotionalCountHandler(c *gin.Context) {
	userIDStr := c.Param("user_id")
	userid, err := strconv.Atoi(userIDStr)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "userID获取失败",
		})
		return
	}
	result, err := logic.GetEmotionalInfo(c.Request.Context(), int64(userid))

	if err != nil {
		zap.L().Error("logic.GetEmotionalInfo failed", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
			"msg": "success",
			"data": result,
		},
	)
}
