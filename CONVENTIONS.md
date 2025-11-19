# VoiceWriter 项目开发规范

本文档定义了 VoiceWriter 项目的代码规范、架构设计和最佳实践。**所有开发者必须严格遵守这些规范。**

---

## 📐 架构设计原则

### 1. 分层架构（Layered Architecture）

项目采用严格的分层架构，各层职责清晰，单向依赖：

```
┌─────────────────────────────────────┐
│   Handler Layer (HTTP处理层)        │  ← 处理HTTP请求/响应
├─────────────────────────────────────┤
│   Service Layer (业务逻辑层)        │  ← 核心业务逻辑
├─────────────────────────────────────┤
│   Repository Layer (数据访问层)     │  ← 数据库操作
├─────────────────────────────────────┤
│   Model Layer (数据模型层)          │  ← 数据结构定义
└─────────────────────────────────────┘
```

**依赖规则：**
- Handler → Service → Repository → Model
- 上层可以调用下层，下层**禁止**调用上层
- 同层之间**禁止**相互调用

### 2. 设计模式

#### Repository Pattern（仓储模式）
- 所有数据库操作必须通过 Repository 接口
- 每个实体对应一个 Repository 接口和实现
- Repository 只负责数据持久化，不包含业务逻辑

#### Dependency Injection（依赖注入）
- 使用接口定义依赖关系
- 通过构造函数注入依赖
- 便于单元测试和 Mock

#### Factory Pattern（工厂模式）
- 复杂对象创建使用工厂方法
- 统一管理对象生命周期

---

## 🗂️ 后端目录结构

```
backend/
├── cmd/
│   └── main.go                 # 应用入口，负责初始化和启动
├── etc/
│   ├── config.yaml            # 默认配置文件
│   ├── config.dev.yaml        # 开发环境配置
│   └── config.prod.yaml       # 生产环境配置
├── internal/
│   ├── config/
│   │   └── config.go          # 配置结构体和加载逻辑
│   ├── model/
│   │   ├── scene.go           # Scene 模型
│   │   ├── sentence.go        # Sentence 模型
│   │   └── user_progress.go   # UserProgress 模型
│   ├── repository/
│   │   ├── interface.go       # Repository 接口定义
│   │   ├── scene_repo.go      # Scene Repository 实现
│   │   ├── sentence_repo.go   # Sentence Repository 实现
│   │   └── progress_repo.go   # Progress Repository 实现
│   ├── service/
│   │   ├── scene_service.go   # Scene 业务逻辑
│   │   ├── sentence_service.go # Sentence 业务逻辑
│   │   └── progress_service.go # Progress 业务逻辑
│   ├── handler/
│   │   ├── scene_handler.go   # Scene HTTP处理器
│   │   ├── sentence_handler.go # Sentence HTTP处理器
│   │   └── progress_handler.go # Progress HTTP处理器
│   ├── middleware/
│   │   ├── cors.go            # CORS中间件
│   │   ├── logger.go          # 日志中间件
│   │   └── recovery.go        # 错误恢复中间件
│   └── database/
│       ├── database.go        # 数据库连接管理
│       └── migrations.go      # 数据库迁移
├── api/
│   └── openapi.yaml           # API 文档（OpenAPI规范）
├── pkg/                       # 可复用的公共包
│   ├── response/
│   │   └── response.go        # 统一响应格式
│   └── errors/
│       └── errors.go          # 自定义错误类型
├── scripts/                   # 工具脚本
│   └── init_db.sql           # 数据库初始化脚本
├── go.mod
├── go.sum
└── README.md
```

---

## 💻 编码规范

### 1. Go 代码规范

#### 命名规范
- **包名**: 小写，简短，有意义（如 `handler`, `service`, `repository`）
- **文件名**: 小写+下划线（如 `scene_service.go`）
- **接口**: 以 `Interface` 或能力命名（如 `SceneRepository`, `Reader`）
- **结构体**: 大驼峰（如 `SceneService`）
- **方法/函数**: 大驼峰（公开）或小驼峰（私有）
- **常量**: 大驼峰或全大写+下划线

#### Model 层规范
```go
// ✅ 正确示例
type Scene struct {
    ID          uint      `gorm:"primarykey" json:"id"`
    Name        string    `gorm:"type:varchar(100);not null" json:"name"`
    Description string    `gorm:"type:text" json:"description"`
    Icon        string    `gorm:"type:varchar(50)" json:"icon"`
    CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
    UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
    DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"` // 软删除
}

// TableName 指定表名
func (Scene) TableName() string {
    return "scenes"
}
```

**规则：**
- 必须使用 GORM 标签定义字段类型和约束
- 统一使用软删除（`gorm.DeletedAt`）
- 时间字段使用 `time.Time`，由 GORM 自动管理
- JSON 标签使用下划线命名

#### Repository 层规范
```go
// ✅ 接口定义（interface.go）
type SceneRepository interface {
    Create(ctx context.Context, scene *model.Scene) error
    GetByID(ctx context.Context, id uint) (*model.Scene, error)
    GetAll(ctx context.Context) ([]*model.Scene, error)
    Update(ctx context.Context, scene *model.Scene) error
    Delete(ctx context.Context, id uint) error
}

// ✅ 实现（scene_repo.go）
type sceneRepository struct {
    db *gorm.DB
}

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
```

**规则：**
- 必须定义接口，接口名为 `实体名 + Repository`
- 实现类首字母小写（私有），通过工厂函数创建
- 所有方法必须接收 `context.Context` 作为第一个参数
- 统一错误处理，将 GORM 错误转换为业务错误
- 使用 `WithContext` 传递上下文

#### Service 层规范
```go
// ✅ 正确示例
type SceneService struct {
    sceneRepo repository.SceneRepository
}

func NewSceneService(sceneRepo repository.SceneRepository) *SceneService {
    return &SceneService{
        sceneRepo: sceneRepo,
    }
}

func (s *SceneService) CreateScene(ctx context.Context, req *CreateSceneRequest) (*model.Scene, error) {
    // 1. 参数验证
    if err := s.validateCreateRequest(req); err != nil {
        return nil, err
    }

    // 2. 业务逻辑处理
    scene := &model.Scene{
        Name:        req.Name,
        Description: req.Description,
        Icon:        req.Icon,
    }

    // 3. 调用 Repository
    if err := s.sceneRepo.Create(ctx, scene); err != nil {
        return nil, err
    }

    return scene, nil
}

func (s *SceneService) validateCreateRequest(req *CreateSceneRequest) error {
    if req.Name == "" {
        return errors.New("name is required")
    }
    return nil
}
```

**规则：**
- Service 通过构造函数注入 Repository 依赖
- 所有方法必须接收 `context.Context` 作为第一个参数
- 负责参数验证、业务逻辑、事务控制
- 不直接操作数据库，通过 Repository 接口调用
- 复杂验证逻辑抽取为私有方法

#### Handler 层规范
```go
// ✅ 正确示例
type SceneHandler struct {
    sceneService *service.SceneService
}

func NewSceneHandler(sceneService *service.SceneService) *SceneHandler {
    return &SceneHandler{
        sceneService: sceneService,
    }
}

func (h *SceneHandler) CreateScene(c *gin.Context) {
    var req service.CreateSceneRequest

    // 1. 参数绑定
    if err := c.ShouldBindJSON(&req); err != nil {
        response.Error(c, http.StatusBadRequest, "Invalid request", err)
        return
    }

    // 2. 调用 Service
    scene, err := h.sceneService.CreateScene(c.Request.Context(), &req)
    if err != nil {
        response.Error(c, http.StatusInternalServerError, "Failed to create scene", err)
        return
    }

    // 3. 返回响应
    response.Success(c, scene)
}
```

**规则：**
- Handler 只负责 HTTP 请求处理，不包含业务逻辑
- 使用统一的响应格式（通过 `pkg/response` 包）
- 参数绑定失败返回 400
- 业务逻辑错误根据类型返回对应状态码
- 传递 `c.Request.Context()` 到 Service 层

### 2. 配置管理规范

#### 配置文件格式（YAML）
```yaml
# etc/config.yaml
server:
  port: 8080
  mode: release  # debug, release, test

database:
  host: localhost
  port: 3306
  username: root
  password: your_password
  dbname: voicewriter
  charset: utf8mb4
  max_idle_conns: 10
  max_open_conns: 100
  conn_max_lifetime: 3600  # 秒

log:
  level: info  # debug, info, warn, error
  format: json  # json, text
  output: stdout  # stdout, file

cors:
  allowed_origins:
    - http://localhost:3000
  allowed_methods:
    - GET
    - POST
    - PUT
    - DELETE
  allowed_headers:
    - Origin
    - Content-Type
    - Authorization
```

#### 配置加载代码
```go
// ✅ 使用 Viper 加载配置
type Config struct {
    Server   ServerConfig   `mapstructure:"server"`
    Database DatabaseConfig `mapstructure:"database"`
    Log      LogConfig      `mapstructure:"log"`
    Cors     CorsConfig     `mapstructure:"cors"`
}

func LoadConfig(configPath string) (*Config, error) {
    viper.SetConfigFile(configPath)
    viper.AutomaticEnv()

    if err := viper.ReadInConfig(); err != nil {
        return nil, err
    }

    var config Config
    if err := viper.Unmarshal(&config); err != nil {
        return nil, err
    }

    return &config, nil
}
```

---

## 🔄 数据库规范

### 1. GORM 使用规范

#### 连接管理
```go
// ✅ 正确示例
func NewDatabase(config *config.DatabaseConfig) (*gorm.DB, error) {
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
        config.Username,
        config.Password,
        config.Host,
        config.Port,
        config.DBName,
        config.Charset,
    )

    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
        Logger: logger.Default.LogMode(logger.Info),
        NamingStrategy: schema.NamingStrategy{
            SingularTable: true, // 使用单数表名
        },
    })
    if err != nil {
        return nil, err
    }

    sqlDB, err := db.DB()
    if err != nil {
        return nil, err
    }

    sqlDB.SetMaxIdleConns(config.MaxIdleConns)
    sqlDB.SetMaxOpenConns(config.MaxOpenConns)
    sqlDB.SetConnMaxLifetime(time.Duration(config.ConnMaxLifetime) * time.Second)

    return db, nil
}
```

#### 迁移管理
```go
// ✅ 自动迁移
func AutoMigrate(db *gorm.DB) error {
    return db.AutoMigrate(
        &model.Scene{},
        &model.Sentence{},
        &model.UserProgress{},
    )
}
```

### 2. 查询规范

```go
// ✅ 使用上下文
db.WithContext(ctx).Find(&users)

// ✅ 预加载关联
db.Preload("Sentences").Find(&scenes)

// ✅ 事务处理
func (s *SceneService) BatchCreate(ctx context.Context, scenes []*model.Scene) error {
    return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
        for _, scene := range scenes {
            if err := tx.Create(scene).Error; err != nil {
                return err
            }
        }
        return nil
    })
}

// ✅ 错误处理
if err := db.First(&user, id).Error; err != nil {
    if errors.Is(err, gorm.ErrRecordNotFound) {
        return nil, ErrUserNotFound
    }
    return nil, err
}
```

---

## 🎯 错误处理规范

### 1. 自定义错误
```go
// pkg/errors/errors.go
var (
    ErrNotFound          = errors.New("resource not found")
    ErrInvalidParameter  = errors.New("invalid parameter")
    ErrUnauthorized      = errors.New("unauthorized")
    ErrInternalServer    = errors.New("internal server error")
)

// 业务错误
type BusinessError struct {
    Code    int
    Message string
    Err     error
}

func (e *BusinessError) Error() string {
    return e.Message
}
```

### 2. 错误处理流程
- Repository: 返回原始错误或转换为业务错误
- Service: 处理业务错误，添加上下文信息
- Handler: 将错误转换为 HTTP 状态码和响应

---

## 📦 统一响应格式

```go
// pkg/response/response.go
type Response struct {
    Code    int         `json:"code"`
    Message string      `json:"message"`
    Data    interface{} `json:"data,omitempty"`
}

func Success(c *gin.Context, data interface{}) {
    c.JSON(http.StatusOK, Response{
        Code:    0,
        Message: "success",
        Data:    data,
    })
}

func Error(c *gin.Context, httpCode int, message string, err error) {
    c.JSON(httpCode, Response{
        Code:    -1,
        Message: message,
    })
}
```

---

## ✅ 代码审查清单

提交代码前，请确保：

- [ ] 遵循分层架构，没有跨层调用
- [ ] 所有 Repository 方法都定义了接口
- [ ] Service 通过依赖注入获取 Repository
- [ ] Handler 只处理 HTTP 请求，不包含业务逻辑
- [ ] 使用 Context 传递请求上下文
- [ ] GORM 模型定义了完整的标签
- [ ] 配置从 YAML 文件读取
- [ ] 错误处理规范统一
- [ ] 代码有必要的注释（特别是导出函数）
- [ ] 遵循 Go 官方代码风格（运行 `gofmt`）

---

## 📚 参考资料

- [Go 代码规范](https://github.com/golang/go/wiki/CodeReviewComments)
- [GORM 官方文档](https://gorm.io/zh_CN/docs/)
- [Gin 框架文档](https://gin-gonic.com/zh-cn/docs/)
- [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)

---

**最后更新**: 2025-11-19
