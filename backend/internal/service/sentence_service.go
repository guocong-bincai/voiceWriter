package service

import (
	"context"
	"errors"

	"voicewriter/internal/model"
	"voicewriter/internal/repository"
)

// SentenceService 句子服务
type SentenceService struct {
	sentenceRepo repository.SentenceRepository
}

// NewSentenceService 创建句子服务实例
func NewSentenceService(sentenceRepo repository.SentenceRepository) *SentenceService {
	return &SentenceService{
		sentenceRepo: sentenceRepo,
	}
}

// GetAllSentences 获取所有句子
func (s *SentenceService) GetAllSentences(ctx context.Context) ([]*model.Sentence, error) {
	return s.sentenceRepo.GetAll(ctx)
}

// GetSentenceByID 根据ID获取句子
func (s *SentenceService) GetSentenceByID(ctx context.Context, id uint) (*model.Sentence, error) {
	if id == 0 {
		return nil, errors.New("invalid sentence id")
	}
	return s.sentenceRepo.GetByID(ctx, id)
}

// GetSentencesBySceneID 根据场景ID获取句子列表
func (s *SentenceService) GetSentencesBySceneID(ctx context.Context, sceneID uint) ([]*model.Sentence, error) {
	if sceneID == 0 {
		return nil, errors.New("invalid scene id")
	}
	return s.sentenceRepo.GetBySceneID(ctx, sceneID)
}

// CreateSentence 创建句子
func (s *SentenceService) CreateSentence(ctx context.Context, sentence *model.Sentence) error {
	if sentence.Content == "" {
		return errors.New("sentence content is required")
	}
	if sentence.SceneID == 0 {
		return errors.New("scene id is required")
	}
	return s.sentenceRepo.Create(ctx, sentence)
}

// UpdateSentence 更新句子
func (s *SentenceService) UpdateSentence(ctx context.Context, sentence *model.Sentence) error {
	if sentence.ID == 0 {
		return errors.New("invalid sentence id")
	}
	if sentence.Content == "" {
		return errors.New("sentence content is required")
	}
	return s.sentenceRepo.Update(ctx, sentence)
}

// DeleteSentence 删除句子
func (s *SentenceService) DeleteSentence(ctx context.Context, id uint) error {
	if id == 0 {
		return errors.New("invalid sentence id")
	}
	return s.sentenceRepo.Delete(ctx, id)
}
