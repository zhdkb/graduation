package controllers

import (
	"graduation/models"
	"net/http"

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

	
}
