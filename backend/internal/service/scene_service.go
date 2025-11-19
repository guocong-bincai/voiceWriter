package service

import (
	"context"
	"errors"

	"voicewriter/internal/model"
	"voicewriter/internal/repository"
)

// SceneService 场景服务
type SceneService struct {
	sceneRepo repository.SceneRepository
}

// NewSceneService 创建场景服务实例
func NewSceneService(sceneRepo repository.SceneRepository) *SceneService {
	return &SceneService{
		sceneRepo: sceneRepo,
	}
}

// GetAllScenes 获取所有场景
func (s *SceneService) GetAllScenes(ctx context.Context) ([]*model.Scene, error) {
	return s.sceneRepo.GetAll(ctx)
}

// GetSceneByID 根据ID获取场景
func (s *SceneService) GetSceneByID(ctx context.Context, id uint) (*model.Scene, error) {
	if id == 0 {
		return nil, errors.New("invalid scene id")
	}
	return s.sceneRepo.GetByID(ctx, id)
}

// CreateScene 创建场景
func (s *SceneService) CreateScene(ctx context.Context, scene *model.Scene) error {
	if scene.Name == "" {
		return errors.New("scene name is required")
	}
	return s.sceneRepo.Create(ctx, scene)
}

// UpdateScene 更新场景
func (s *SceneService) UpdateScene(ctx context.Context, scene *model.Scene) error {
	if scene.ID == 0 {
		return errors.New("invalid scene id")
	}
	if scene.Name == "" {
		return errors.New("scene name is required")
	}
	return s.sceneRepo.Update(ctx, scene)
}

// DeleteScene 删除场景
func (s *SceneService) DeleteScene(ctx context.Context, id uint) error {
	if id == 0 {
		return errors.New("invalid scene id")
	}
	return s.sceneRepo.Delete(ctx, id)
}
