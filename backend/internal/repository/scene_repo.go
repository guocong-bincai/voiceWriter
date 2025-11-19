package repository

import (
	"context"
	"errors"

	"voicewriter/internal/model"

	"gorm.io/gorm"
)

type sceneRepository struct {
	db *gorm.DB
}

// NewSceneRepository 创建场景仓储实例
func NewSceneRepository(db *gorm.DB) SceneRepository {
	return &sceneRepository{db: db}
}

func (r *sceneRepository) Create(ctx context.Context, scene *model.Scene) error {
	return r.db.WithContext(ctx).Create(scene).Error
}

func (r *sceneRepository) GetByID(ctx context.Context, id uint) (*model.Scene, error) {
	var scene model.Scene
	err := r.db.WithContext(ctx).First(&scene, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &scene, nil
}

func (r *sceneRepository) GetAll(ctx context.Context) ([]*model.Scene, error) {
	var scenes []*model.Scene
	err := r.db.WithContext(ctx).Find(&scenes).Error
	if err != nil {
		return nil, err
	}
	return scenes, nil
}

func (r *sceneRepository) Update(ctx context.Context, scene *model.Scene) error {
	return r.db.WithContext(ctx).Save(scene).Error
}

func (r *sceneRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.Scene{}, id).Error
}
