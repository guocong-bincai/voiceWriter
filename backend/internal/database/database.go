package database

import (
	"fmt"
	"log"
	"time"

	"voicewriter/internal/config"
	"voicewriter/internal/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// NewDatabase 创建数据库连接
func NewDatabase(cfg *config.DatabaseConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DBName,
		cfg.Charset,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}

	// 获取底层 sql.DB 配置连接池
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get database instance: %w", err)
	}

	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(cfg.ConnMaxLifetime) * time.Second)

	log.Println("Database connection established successfully")
	return db, nil
}

// AutoMigrate 自动迁移数据库表结构
func AutoMigrate(db *gorm.DB) error {
	log.Println("Starting database migration...")

	err := db.AutoMigrate(
		&model.Scene{},
		&model.Sentence{},
		&model.UserProgress{},
	)

	if err != nil {
		return fmt.Errorf("failed to migrate database: %w", err)
	}

	log.Println("Database migration completed successfully")
	return nil
}

// SeedData 初始化种子数据
func SeedData(db *gorm.DB) error {
	// 检查是否已有数据
	var count int64
	db.Model(&model.Scene{}).Count(&count)
	if count > 0 {
		log.Println("Database already has data, skipping seed")
		return nil
	}

	log.Println("Seeding initial data...")

	// 创建场景数据
	scenes := []*model.Scene{
		{
			Name:        "日常生活",
			Description: "日常生活中的常用对话",
			Icon:        "home",
		},
		{
			Name:        "工作职场",
			Description: "工作场景中的专业对话",
			Icon:        "work",
		},
		{
			Name:        "旅游出行",
			Description: "旅游时的实用对话",
			Icon:        "travel",
		},
	}

	for _, scene := range scenes {
		if err := db.Create(scene).Error; err != nil {
			return fmt.Errorf("failed to create scene: %w", err)
		}
	}

	// 创建句子数据
	sentences := []*model.Sentence{
		{
			SceneID:     1,
			Content:     "Hello, how are you?",
			Translation: "你好，你怎么样？",
			AudioURL:    "/audio/1.mp3",
			Difficulty:  "easy",
		},
		{
			SceneID:     1,
			Content:     "What's your name?",
			Translation: "你叫什么名字？",
			AudioURL:    "/audio/2.mp3",
			Difficulty:  "easy",
		},
		{
			SceneID:     1,
			Content:     "Nice to meet you!",
			Translation: "很高兴见到你！",
			AudioURL:    "/audio/3.mp3",
			Difficulty:  "easy",
		},
		{
			SceneID:     2,
			Content:     "Could you please send me the report?",
			Translation: "你能把报告发给我吗？",
			AudioURL:    "/audio/4.mp3",
			Difficulty:  "medium",
		},
		{
			SceneID:     2,
			Content:     "Let's schedule a meeting for next week.",
			Translation: "我们下周安排一个会议吧。",
			AudioURL:    "/audio/5.mp3",
			Difficulty:  "medium",
		},
		{
			SceneID:     3,
			Content:     "How much does this cost?",
			Translation: "这个多少钱？",
			AudioURL:    "/audio/6.mp3",
			Difficulty:  "easy",
		},
		{
			SceneID:     3,
			Content:     "Where is the nearest subway station?",
			Translation: "最近的地铁站在哪里？",
			AudioURL:    "/audio/7.mp3",
			Difficulty:  "medium",
		},
	}

	for _, sentence := range sentences {
		if err := db.Create(sentence).Error; err != nil {
			return fmt.Errorf("failed to create sentence: %w", err)
		}
	}

	log.Println("Seed data created successfully")
	return nil
}
