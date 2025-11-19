package handler

import (
	"strconv"

	"voicewriter/internal/service"
	"voicewriter/pkg/response"

	"github.com/gin-gonic/gin"
)

// SentenceHandler 句子处理器
type SentenceHandler struct {
	sentenceService *service.SentenceService
}

// NewSentenceHandler 创建句子处理器实例
func NewSentenceHandler(sentenceService *service.SentenceService) *SentenceHandler {
	return &SentenceHandler{
		sentenceService: sentenceService,
	}
}

// GetSentences 获取所有句子
// @Summary 获取所有句子
// @Description 获取所有句子列表
// @Tags 句子
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Router /api/v1/sentences [get]
func (h *SentenceHandler) GetSentences(c *gin.Context) {
	sentences, err := h.sentenceService.GetAllSentences(c.Request.Context())
	if err != nil {
		response.InternalServerError(c, "Failed to get sentences")
		return
	}

	response.Success(c, sentences)
}

// GetSentenceByID 根据ID获取句子
// @Summary 根据ID获取句子
// @Description 根据句子ID获取句子详情
// @Tags 句子
// @Accept json
// @Produce json
// @Param id path int true "句子ID"
// @Success 200 {object} response.Response
// @Router /api/v1/sentences/{id} [get]
func (h *SentenceHandler) GetSentenceByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid sentence ID")
		return
	}

	sentence, err := h.sentenceService.GetSentenceByID(c.Request.Context(), uint(id))
	if err != nil {
		response.NotFound(c, "Sentence not found")
		return
	}

	response.Success(c, sentence)
}

// GetSentencesByScene 根据场景ID获取句子列表
// @Summary 根据场景ID获取句子列表
// @Description 获取指定场景下的所有句子
// @Tags 句子
// @Accept json
// @Produce json
// @Param sceneId path int true "场景ID"
// @Success 200 {object} response.Response
// @Router /api/v1/sentences/scene/{sceneId} [get]
func (h *SentenceHandler) GetSentencesByScene(c *gin.Context) {
	sceneIDStr := c.Param("sceneId")
	sceneID, err := strconv.ParseUint(sceneIDStr, 10, 32)
	if err != nil {
		response.BadRequest(c, "Invalid scene ID")
		return
	}

	sentences, err := h.sentenceService.GetSentencesBySceneID(c.Request.Context(), uint(sceneID))
	if err != nil {
		response.InternalServerError(c, "Failed to get sentences")
		return
	}

	response.Success(c, sentences)
}
