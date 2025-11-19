package handler

import (
	"voicewriter/internal/model"
	"voicewriter/internal/service"
	"voicewriter/pkg/response"

	"github.com/gin-gonic/gin"
)

// ProgressHandler 用户进度处理器
type ProgressHandler struct {
	progressService *service.ProgressService
}

// NewProgressHandler 创建用户进度处理器实例
func NewProgressHandler(progressService *service.ProgressService) *ProgressHandler {
	return &ProgressHandler{
		progressService: progressService,
	}
}

// GetUserProgress 获取用户进度
// @Summary 获取用户进度
// @Description 根据用户ID获取学习进度
// @Tags 进度
// @Accept json
// @Produce json
// @Param userId path string true "用户ID"
// @Success 200 {object} response.Response
// @Router /api/v1/progress/{userId} [get]
func (h *ProgressHandler) GetUserProgress(c *gin.Context) {
	userID := c.Param("userId")
	if userID == "" {
		response.BadRequest(c, "User ID is required")
		return
	}

	progress, err := h.progressService.GetUserProgress(c.Request.Context(), userID)
	if err != nil {
		response.InternalServerError(c, "Failed to get user progress")
		return
	}

	response.Success(c, progress)
}

// SaveUserProgress 保存用户进度
// @Summary 保存用户进度
// @Description 保存或更新用户学习进度
// @Tags 进度
// @Accept json
// @Produce json
// @Param progress body model.UserProgress true "进度信息"
// @Success 200 {object} response.Response
// @Router /api/v1/progress [post]
func (h *ProgressHandler) SaveUserProgress(c *gin.Context) {
	var progress model.UserProgress
	if err := c.ShouldBindJSON(&progress); err != nil {
		response.BadRequest(c, "Invalid request body")
		return
	}

	if err := h.progressService.SaveProgress(c.Request.Context(), &progress); err != nil {
		response.InternalServerError(c, "Failed to save progress")
		return
	}

	response.SuccessWithMessage(c, "Progress saved successfully", nil)
}
