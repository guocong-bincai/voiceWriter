package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthCheck 健康检查
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "VoiceWriter API is running",
	})
}

// GetAudio 获取音频
// TODO: 实现音频获取逻辑（集成TTS或返回预录音频）
func GetAudio(c *gin.Context) {
	id := c.Param("id")

	// 暂时返回占位符
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": gin.H{
			"url": "/audio/" + id + ".mp3",
		},
		"message": "success",
	})
}
