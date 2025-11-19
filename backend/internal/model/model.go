package model

import (
	"time"

	"gorm.io/gorm"
)

// Scene 场景模型
type Scene struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	Name        string         `gorm:"type:varchar(100);not null" json:"name"`
	Description string         `gorm:"type:text" json:"description"`
	Icon        string         `gorm:"type:varchar(50)" json:"icon"`
	CreatedAt   time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`

	// 关联
	Sentences []Sentence `gorm:"foreignKey:SceneID" json:"sentences,omitempty"`
}

// TableName 指定表名
func (Scene) TableName() string {
	return "scenes"
}

// Sentence 句子模型
type Sentence struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	SceneID     uint           `gorm:"not null;index" json:"scene_id"`
	Content     string         `gorm:"type:text;not null" json:"content"`
	Translation string         `gorm:"type:text" json:"translation"`
	AudioURL    string         `gorm:"type:varchar(255)" json:"audio_url"`
	Difficulty  string         `gorm:"type:varchar(20);default:'easy'" json:"difficulty"` // easy, medium, hard
	CreatedAt   time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`

	// 关联
	Scene *Scene `gorm:"foreignKey:SceneID" json:"scene,omitempty"`
}

// TableName 指定表名
func (Sentence) TableName() string {
	return "sentences"
}

// UserProgress 用户进度模型
type UserProgress struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	UserID      string         `gorm:"type:varchar(100);not null;index" json:"user_id"`
	SentenceID  uint           `gorm:"not null;index" json:"sentence_id"`
	Completed   bool           `gorm:"default:false" json:"completed"`
	Attempts    int            `gorm:"default:0" json:"attempts"`
	LastAttempt time.Time      `json:"last_attempt"`
	CreatedAt   time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`

	// 关联
	Sentence *Sentence `gorm:"foreignKey:SentenceID" json:"sentence,omitempty"`
}

// TableName 指定表名
func (UserProgress) TableName() string {
	return "user_progress"
}
