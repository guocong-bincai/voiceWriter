package repository

import (
	"context"
	"errors"

	"voicewriter/internal/model"

	"gorm.io/gorm"
)

type progressRepository struct {
	db *gorm.DB
}

// NewProgressRepository 创建用户进度仓储实例
func NewProgressRepository(db *gorm.DB) ProgressRepository {
	return &progressRepository{db: db}
}

func (r *progressRepository) Create(ctx context.Context, progress *model.UserProgress) error {
	return r.db.WithContext(ctx).Create(progress).Error
}

func (r *progressRepository) GetByID(ctx context.Context, id uint) (*model.UserProgress, error) {
	var progress model.UserProgress
	err := r.db.WithContext(ctx).First(&progress, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &progress, nil
}

func (r *progressRepository) GetByUserID(ctx context.Context, userID string) ([]*model.UserProgress, error) {
	var progress []*model.UserProgress
	err := r.db.WithContext(ctx).Where("user_id = ?", userID).Find(&progress).Error
	if err != nil {
		return nil, err
	}
	return progress, nil
}

func (r *progressRepository) GetByUserAndSentence(ctx context.Context, userID string, sentenceID uint) (*model.UserProgress, error) {
	var progress model.UserProgress
	err := r.db.WithContext(ctx).
		Where("user_id = ? AND sentence_id = ?", userID, sentenceID).
		First(&progress).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &progress, nil
}

func (r *progressRepository) Update(ctx context.Context, progress *model.UserProgress) error {
	return r.db.WithContext(ctx).Save(progress).Error
}

func (r *progressRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.UserProgress{}, id).Error
}
