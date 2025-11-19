package main

import (
	"log"
	"os"

	"voicewriter/internal/config"
	"voicewriter/internal/database"
	"voicewriter/internal/handler"
	"voicewriter/internal/repository"
	"voicewriter/internal/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// 加载配置
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		configPath = "etc/config.yaml"
	}

	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// 设置Gin模式
	gin.SetMode(cfg.Server.Mode)

	// 初始化数据库连接
	db, err := database.NewDatabase(&cfg.Database)
	if err != nil {
		log.Fatalf("Failed to connect database: %v", err)
	}

	// 自动迁移数据库表结构
	if err := database.AutoMigrate(db); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// 初始化种子数据
	if err := database.SeedData(db); err != nil {
		log.Fatalf("Failed to seed data: %v", err)
	}

	// 初始化Repository层
	sceneRepo := repository.NewSceneRepository(db)
	sentenceRepo := repository.NewSentenceRepository(db)
	progressRepo := repository.NewProgressRepository(db)

	// 初始化Service层
	sceneService := service.NewSceneService(sceneRepo)
	sentenceService := service.NewSentenceService(sentenceRepo)
	progressService := service.NewProgressService(progressRepo)

	// 初始化Handler层
	sceneHandler := handler.NewSceneHandler(sceneService)
	sentenceHandler := handler.NewSentenceHandler(sentenceService)
	progressHandler := handler.NewProgressHandler(progressService)

	// 创建Gin引擎
	r := gin.Default()

	// 配置CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     cfg.Cors.AllowedOrigins,
		AllowMethods:     cfg.Cors.AllowedMethods,
		AllowHeaders:     cfg.Cors.AllowedHeaders,
		AllowCredentials: true,
	}))

	// 注册路由
	setupRoutes(r, sceneHandler, sentenceHandler, progressHandler)

	// 启动服务
	addr := ":" + cfg.Server.Port
	log.Printf("Server starting on %s in %s mode", addr, cfg.Server.Mode)
	if err := r.Run(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func setupRoutes(
	r *gin.Engine,
	sceneHandler *handler.SceneHandler,
	sentenceHandler *handler.SentenceHandler,
	progressHandler *handler.ProgressHandler,
) {
	// 健康检查
	r.GET("/health", handler.HealthCheck)

	// API v1
	v1 := r.Group("/api/v1")
	{
		// 场景相关
		scenes := v1.Group("/scenes")
		{
			scenes.GET("", sceneHandler.GetScenes)
			scenes.GET("/:id", sceneHandler.GetSceneByID)
		}

		// 句子相关
		sentences := v1.Group("/sentences")
		{
			sentences.GET("", sentenceHandler.GetSentences)
			sentences.GET("/:id", sentenceHandler.GetSentenceByID)
			sentences.GET("/scene/:sceneId", sentenceHandler.GetSentencesByScene)
		}

		// 音频相关
		audio := v1.Group("/audio")
		{
			audio.GET("/:id", handler.GetAudio)
		}

		// 用户进度相关
		progress := v1.Group("/progress")
		{
			progress.GET("/:userId", progressHandler.GetUserProgress)
			progress.POST("", progressHandler.SaveUserProgress)
		}
	}
}
