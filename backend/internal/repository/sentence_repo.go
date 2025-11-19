package repository

import (
	"context"
	"errors"

	"voicewriter/internal/model"

	"gorm.io/gorm"
)

type sentenceRepository struct {
	db *gorm.DB
}

// NewSentenceRepository 创建句子仓储实例
func NewSentenceRepository(db *gorm.DB) SentenceRepository {
	return &sentenceRepository{db: db}
}

func (r *sentenceRepository) Create(ctx context.Context, sentence *model.Sentence) error {
	return r.db.WithContext(ctx).Create(sentence).Error
}

func (r *sentenceRepository) GetByID(ctx context.Context, id uint) (*model.Sentence, error) {
	var sentence model.Sentence
	err := r.db.WithContext(ctx).First(&sentence, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &sentence, nil
}

func (r *sentenceRepository) GetAll(ctx context.Context) ([]*model.Sentence, error) {
	var sentences []*model.Sentence
	err := r.db.WithContext(ctx).Find(&sentences).Error
	if err != nil {
		return nil, err
	}
	return sentences, nil
}

func (r *sentenceRepository) GetBySceneID(ctx context.Context, sceneID uint) ([]*model.Sentence, error) {
	var sentences []*model.Sentence
	err := r.db.WithContext(ctx).Where("scene_id = ?", sceneID).Find(&sentences).Error
	if err != nil {
		return nil, err
	}
	return sentences, nil
}

func (r *sentenceRepository) Update(ctx context.Context, sentence *model.Sentence) error {
	return r.db.WithContext(ctx).Save(sentence).Error
}

func (r *sentenceRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.Sentence{}, id).Error
}
