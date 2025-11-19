package repository

import (
	"context"
	"errors"

	"voicewriter/internal/model"
)

var (
	// ErrNotFound 记录未找到错误
	ErrNotFound = errors.New("record not found")
	// ErrDuplicateKey 唯一键冲突错误
	ErrDuplicateKey = errors.New("duplicate key")
)

// SceneRepository 场景仓储接口
type SceneRepository interface {
	Create(ctx context.Context, scene *model.Scene) error
	GetByID(ctx context.Context, id uint) (*model.Scene, error)
	GetAll(ctx context.Context) ([]*model.Scene, error)
	Update(ctx context.Context, scene *model.Scene) error
	Delete(ctx context.Context, id uint) error
}

// SentenceRepository 句子仓储接口
type SentenceRepository interface {
	Create(ctx context.Context, sentence *model.Sentence) error
	GetByID(ctx context.Context, id uint) (*model.Sentence, error)
	GetAll(ctx context.Context) ([]*model.Sentence, error)
	GetBySceneID(ctx context.Context, sceneID uint) ([]*model.Sentence, error)
	Update(ctx context.Context, sentence *model.Sentence) error
	Delete(ctx context.Context, id uint) error
}

// ProgressRepository 用户进度仓储接口
type ProgressRepository interface {
	Create(ctx context.Context, progress *model.UserProgress) error
	GetByID(ctx context.Context, id uint) (*model.UserProgress, error)
	GetByUserID(ctx context.Context, userID string) ([]*model.UserProgress, error)
	GetByUserAndSentence(ctx context.Context, userID string, sentenceID uint) (*model.UserProgress, error)
	Update(ctx context.Context, progress *model.UserProgress) error
	Delete(ctx context.Context, id uint) error
}
