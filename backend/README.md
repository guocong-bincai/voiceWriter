# VoiceWriter Backend

基于 Gin + GORM + MySQL 的英语听写练习后端服务

## 技术栈

- **语言**: Go 1.21+
- **Web框架**: Gin
- **ORM**: GORM
- **数据库**: MySQL 8.0+
- **配置管理**: Viper (YAML)

## 项目结构

```
backend/
├── cmd/
│   └── main.go              # 应用程序入口
├── etc/
│   └── config.yaml          # 配置文件
├── internal/                # 内部包
│   ├── config/              # 配置管理
│   ├── database/            # 数据库连接和迁移
│   ├── model/               # 数据模型
│   ├── repository/          # 数据访问层（Repository Pattern）
│   ├── service/             # 业务逻辑层
│   ├── handler/             # HTTP处理层
│   └── middleware/          # 中间件
├── pkg/                     # 可复用的公共包
│   ├── response/            # 统一响应格式
│   └── errors/              # 自定义错误
├── scripts/
│   └── init_db.sql          # 数据库初始化脚本
├── go.mod
└── README.md
```

## 架构设计

项目采用严格的分层架构：

```
Handler Layer (HTTP处理)
    ↓
Service Layer (业务逻辑)
    ↓
Repository Layer (数据访问)
    ↓
Model Layer (数据模型)
```

**核心原则**:
- 依赖注入 (Dependency Injection)
- 仓储模式 (Repository Pattern)
- 接口驱动开发
- 单向依赖（上层依赖下层）

详细规范请查看 [CONVENTIONS.md](../CONVENTIONS.md)

## 快速开始

### 1. 安装依赖

```bash
cd backend
go mod download
```

### 2. 配置数据库

编辑 `etc/config.yaml`:

```yaml
database:
  host: localhost
  port: 3306
  username: root
  password: your_password
  dbname: voicewriter
  charset: utf8mb4
```

或创建数据库:

```bash
mysql -u root -p < scripts/init_db.sql
```

### 3. 运行服务

```bash
# 使用默认配置
go run cmd/main.go

# 或指定配置文件
CONFIG_PATH=etc/config.yaml go run cmd/main.go
```

服务将在 `http://localhost:8080` 启动

### 4. 验证服务

```bash
curl http://localhost:8080/health
```

## API 接口

### 场景管理
- `GET /api/v1/scenes` - 获取所有场景
- `GET /api/v1/scenes/:id` - 获取指定场景

### 句子管理
- `GET /api/v1/sentences` - 获取所有句子
- `GET /api/v1/sentences/:id` - 获取指定句子
- `GET /api/v1/sentences/scene/:sceneId` - 获取场景下的句子

### 音频管理
- `GET /api/v1/audio/:id` - 获取音频文件

### 用户进度
- `GET /api/v1/progress/:userId` - 获取用户进度
- `POST /api/v1/progress` - 保存用户进度

## 配置说明

配置文件 `etc/config.yaml` 包含以下配置项：

```yaml
server:
  port: 8080                # 服务端口
  mode: debug               # 运行模式: debug, release, test

database:
  host: localhost           # 数据库主机
  port: 3306               # 数据库端口
  username: root           # 数据库用户名
  password: your_password  # 数据库密码
  dbname: voicewriter      # 数据库名称
  charset: utf8mb4         # 字符集
  max_idle_conns: 10       # 最大空闲连接数
  max_open_conns: 100      # 最大打开连接数
  conn_max_lifetime: 3600  # 连接最大生命周期(秒)

log:
  level: info              # 日志级别
  format: json             # 日志格式
  output: stdout           # 日志输出

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

## 数据库设计

### scenes (场景表)
| 字段 | 类型 | 说明 |
|------|------|------|
| id | INT UNSIGNED | 主键 |
| name | VARCHAR(100) | 场景名称 |
| description | TEXT | 场景描述 |
| icon | VARCHAR(50) | 图标 |
| created_at | TIMESTAMP | 创建时间 |
| updated_at | TIMESTAMP | 更新时间 |
| deleted_at | TIMESTAMP | 删除时间（软删除） |

### sentences (句子表)
| 字段 | 类型 | 说明 |
|------|------|------|
| id | INT UNSIGNED | 主键 |
| scene_id | INT UNSIGNED | 场景ID（外键） |
| content | TEXT | 英文句子 |
| translation | TEXT | 中文翻译 |
| audio_url | VARCHAR(255) | 音频URL |
| difficulty | VARCHAR(20) | 难度：easy, medium, hard |
| created_at | TIMESTAMP | 创建时间 |
| updated_at | TIMESTAMP | 更新时间 |
| deleted_at | TIMESTAMP | 删除时间（软删除） |

### user_progress (用户进度表)
| 字段 | 类型 | 说明 |
|------|------|------|
| id | INT UNSIGNED | 主键 |
| user_id | VARCHAR(100) | 用户ID |
| sentence_id | INT UNSIGNED | 句子ID（外键） |
| completed | BOOLEAN | 是否完成 |
| attempts | INT | 尝试次数 |
| last_attempt | TIMESTAMP | 最后尝试时间 |
| created_at | TIMESTAMP | 创建时间 |
| updated_at | TIMESTAMP | 更新时间 |
| deleted_at | TIMESTAMP | 删除时间（软删除） |

## 开发指南

### 添加新功能

1. **定义Model** (internal/model/)
2. **创建Repository接口和实现** (internal/repository/)
3. **实现Service业务逻辑** (internal/service/)
4. **创建Handler处理HTTP请求** (internal/handler/)
5. **在main.go中注册路由**

### 代码规范

- 遵循 [CONVENTIONS.md](../CONVENTIONS.md) 中的规范
- 运行 `gofmt` 格式化代码
- 所有导出函数必须有注释
- 使用依赖注入
- 错误处理要规范

### 测试

```bash
# 运行测试
go test ./...

# 运行测试并生成覆盖率报告
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

## 待实现功能

- [ ] TTS 音频服务集成
- [ ] 用户认证和授权
- [ ] 单元测试和集成测试
- [ ] API 文档自动生成（Swagger）
- [ ] 日志中间件
- [ ] 限流中间件
- [ ] 缓存支持（Redis）
- [ ] Docker 部署

## 故障排查

### 数据库连接失败
- 检查 MySQL 是否运行
- 验证 config.yaml 中的数据库配置
- 确认数据库已创建

### 端口被占用
- 修改 config.yaml 中的 port 配置
- 或使用环境变量 `PORT=8081 go run cmd/main.go`

## 许可证

MIT
