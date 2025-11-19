package service

import (
	"context"
	"errors"
	"time"

	"voicewriter/internal/model"
	"voicewriter/internal/repository"
)

// ProgressService 用户进度服务
type ProgressService struct {
	progressRepo repository.ProgressRepository
}

// NewProgressService 创建用户进度服务实例
func NewProgressService(progressRepo repository.ProgressRepository) *ProgressService {
	return &ProgressService{
		progressRepo: progressRepo,
	}
}

// GetUserProgress 获取用户进度
func (s *ProgressService) GetUserProgress(ctx context.Context, userID string) ([]*model.UserProgress, error) {
	if userID == "" {
		return nil, errors.New("user id is required")
	}
	return s.progressRepo.GetByUserID(ctx, userID)
}

// SaveProgress 保存用户进度
func (s *ProgressService) SaveProgress(ctx context.Context, progress *model.UserProgress) error {
	if progress.UserID == "" {
		return errors.New("user id is required")
	}
	if progress.SentenceID == 0 {
		return errors.New("sentence id is required")
	}

	// 检查是否已存在
	existing, err := s.progressRepo.GetByUserAndSentence(ctx, progress.UserID, progress.SentenceID)
	if err != nil && !errors.Is(err, repository.ErrNotFound) {
		return err
	}

	// 如果已存在，更新记录
	if existing != nil {
		existing.Completed = progress.Completed
		existing.Attempts++
		existing.LastAttempt = time.Now()
		return s.progressRepo.Update(ctx, existing)
	}

	// 否则创建新记录
	progress.Attempts = 1
	progress.LastAttempt = time.Now()
	return s.progressRepo.Create(ctx, progress)
}

// GetProgressByID 根据ID获取进度
func (s *ProgressService) GetProgressByID(ctx context.Context, id uint) (*model.UserProgress, error) {
	if id == 0 {
		return nil, errors.New("invalid progress id")
	}
	return s.progressRepo.GetByID(ctx, id)
}

// DeleteProgress 删除进度
func (s *ProgressService) DeleteProgress(ctx context.Context, id uint) error {
	if id == 0 {
		return errors.New("invalid progress id")
	}
	return s.progressRepo.Delete(ctx, id)
}
