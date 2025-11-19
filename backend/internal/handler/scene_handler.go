package handler

import (
	"net/http"
	"strconv"

	"voicewriter/internal/service"
	"voicewriter/pkg/response"

	"github.com/gin-gonic/gin"
)

// SceneHandler 场景处理器
type SceneHandler struct {
	sceneService *service.SceneService
}

// NewSceneHandler 创建场景处理器实例
func NewSceneHandler(sceneService *service.SceneService) *SceneHandler {
	return &SceneHandler{
		sceneService: sceneService,
	}
}

// GetScenes 获取所有场景
// @Summary 获取所有场景
// @Description 获取所有场景列表
// @Tags 场景
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Router /api/v1/scenes [get]
func (h *SceneHandler) GetScenes(c *gin.Context) {
	scenes, err := h.sceneService.GetAllScenes(c.Request.Context())
	if err != nil {
		response.InternalServerError(c, "Failed to get scenes")
		return
	}

	response.Success(c, scenes)
}

// GetSceneByID 根据ID获取场景
// @Summary 根据ID获取场景
// @Description 根据场景ID获取场景详情
// @Tags 场景
// @Accept json
// @Produce json
// @Param id path int true "场景ID"
// @Success 200 {object} response.Response
// @Router /api/v1/scenes/{id} [get]
func (h *SceneHandler) GetSceneByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid scene ID")
		return
	}

	scene, err := h.sceneService.GetSceneByID(c.Request.Context(), uint(id))
	if err != nil {
		response.NotFound(c, "Scene not found")
		return
	}

	response.Success(c, scene)
}
