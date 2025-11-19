# VoiceWriter - 多语言听写练习平台

[English](#english) | [中文](#中文) | [한국어](#한국어) | [日本語](#日本語)

---

## 中文

一个帮助用户通过听写练习提高外语水平的Web应用。通过真实生活场景的对话句子，让用户在听音默写的过程中发现并记忆陌生单词。

### 🌍 支持的语言

- **🇨🇳 中文** - 学习英语、日语、韩语
- **🇺🇸 英语** - 学习中文、日语、韩语
- **🇰🇷 韩语** - 学习英语、中文、日语
- **🇯🇵 日语** - 学习英语、中文、韩语

### 🎯 核心功能

- **📚 场景化学习**: 提供日常生活、工作职场、旅游出行等多种真实场景
- **🎧 听音默写**: 点击播放音频，听写目标语言句子
- **⚡ 即时反馈**: 提交后立即显示正确答案和翻译
- **📊 难度分级**: 句子按照难度分为简单、中等、困难三个级别
- **📈 进度追踪**: 记录用户的学习进度和成果
- **🌐 多语言界面**: 支持中文、英文、韩文、日文界面切换

## 🏗️ 项目架构

```
VoiceWriter/
├── backend/          # Go后端服务
│   ├── cmd/         # 应用入口
│   ├── internal/    # 内部包
│   │   ├── config/  # 配置管理
│   │   ├── handler/ # HTTP处理器
│   │   ├── model/   # 数据模型
│   │   ├── service/ # 业务逻辑
│   │   └── repository/ # 数据访问
│   └── api/         # API文档
└── frontend/        # React前端应用
    ├── src/
    │   ├── components/ # UI组件
    │   ├── pages/      # 页面
    │   ├── services/   # API服务
    │   ├── store/      # 状态管理
    │   └── hooks/      # 自定义Hooks
    └── public/         # 静态资源
```

## 🛠️ 技术栈

### 后端
- **语言**: Go 1.21+
- **框架**: Gin
- **数据库**: PostgreSQL (计划)
- **音频**: TTS服务集成 (计划)

### 前端
- **框架**: React 18 + TypeScript
- **UI库**: Ant Design
- **样式**: Tailwind CSS
- **状态管理**: Zustand
- **路由**: React Router
- **HTTP**: Axios
- **音频**: Howler.js

## 🚀 快速开始

### 后端服务

```bash
cd backend
go mod download
go run cmd/main.go
```

后端服务将在 `http://localhost:8080` 启动

### 前端应用

```bash
cd frontend
npm install
npm start
```

前端应用将在 `http://localhost:3000` 启动

## 📝 API文档

### 场景相关
- `GET /api/v1/scenes` - 获取所有场景
- `GET /api/v1/scenes/:id` - 获取指定场景

### 句子相关
- `GET /api/v1/sentences` - 获取所有句子
- `GET /api/v1/sentences/:id` - 获取指定句子
- `GET /api/v1/sentences/scene/:sceneId` - 获取场景下的句子

### 音频相关
- `GET /api/v1/audio/:id` - 获取音频文件

### 用户进度
- `GET /api/v1/progress/:userId` - 获取用户进度
- `POST /api/v1/progress` - 保存用户进度

### 📅 开发计划

- [x] 前后端框架搭建
- [x] 基础API接口
- [x] 页面布局和路由
- [x] 音频播放功能
- [x] 数据库集成（MySQL + GORM）
- [x] 严格分层架构实现
- [x] YAML配置管理
- [ ] 多语言TTS服务集成
- [ ] 用户认证系统
- [ ] 进度统计分析
- [ ] 多语言界面支持
- [ ] 移动端适配
- [ ] 单元测试和集成测试

### 📄 许可证

MIT License

### 👥 贡献

欢迎提交Issue和Pull Request！

---

## English

A web application that helps users improve their foreign language skills through dictation practice. Through real-life scenario dialogues, users can discover and memorize unfamiliar words during the listening and writing process.

### 🌍 Supported Languages

- **🇨🇳 Chinese** - Learn English, Japanese, Korean
- **🇺🇸 English** - Learn Chinese, Japanese, Korean
- **🇰🇷 Korean** - Learn English, Chinese, Japanese
- **🇯🇵 Japanese** - Learn English, Chinese, Korean

### 🎯 Core Features

- **📚 Scenario-based Learning**: Various real-life scenarios including daily life, workplace, and travel
- **🎧 Audio Dictation**: Click to play audio and write down what you hear
- **⚡ Instant Feedback**: Immediate display of correct answers and translations
- **📊 Difficulty Levels**: Sentences categorized into easy, medium, and hard levels
- **📈 Progress Tracking**: Record user learning progress and achievements
- **🌐 Multilingual Interface**: Support for Chinese, English, Korean, and Japanese interfaces

### 🏗️ Project Architecture

```
VoiceWriter/
├── backend/          # Go backend service
│   ├── cmd/         # Application entry
│   ├── internal/    # Internal packages
│   │   ├── config/  # Configuration management
│   │   ├── handler/ # HTTP handlers
│   │   ├── model/   # Data models
│   │   ├── service/ # Business logic
│   │   └── repository/ # Data access
│   └── api/         # API documentation
└── frontend/        # React frontend application
    ├── src/
    │   ├── components/ # UI components
    │   ├── pages/      # Pages
    │   ├── services/   # API services
    │   ├── store/      # State management
    │   └── hooks/      # Custom hooks
    └── public/         # Static assets
```

### 🛠️ Tech Stack

#### Backend
- **Language**: Go 1.21+
- **Framework**: Gin
- **Database**: MySQL + GORM
- **Config**: Viper (YAML)

#### Frontend
- **Framework**: React 18 + TypeScript
- **UI Library**: Ant Design
- **Styling**: Tailwind CSS
- **State Management**: Zustand
- **Router**: React Router
- **HTTP**: Axios
- **Audio**: Howler.js

### 🚀 Quick Start

#### Backend Service

```bash
cd backend
go mod download
go run cmd/main.go
```

Backend service will start at `http://localhost:8080`

#### Frontend Application

```bash
cd frontend
npm install
npm start
```

Frontend application will start at `http://localhost:3000`

### 📝 API Documentation

#### Scene Management
- `GET /api/v1/scenes` - Get all scenes
- `GET /api/v1/scenes/:id` - Get scene by ID

#### Sentence Management
- `GET /api/v1/sentences` - Get all sentences
- `GET /api/v1/sentences/:id` - Get sentence by ID
- `GET /api/v1/sentences/scene/:sceneId` - Get sentences by scene

#### Audio Management
- `GET /api/v1/audio/:id` - Get audio file

#### User Progress
- `GET /api/v1/progress/:userId` - Get user progress
- `POST /api/v1/progress` - Save user progress

### 📅 Development Roadmap

- [x] Frontend and backend framework setup
- [x] Basic API endpoints
- [x] Page layout and routing
- [x] Audio playback functionality
- [x] Database integration (MySQL + GORM)
- [x] Strict layered architecture implementation
- [x] YAML configuration management
- [ ] Multilingual TTS service integration
- [ ] User authentication system
- [ ] Progress statistics and analysis
- [ ] Multilingual interface support
- [ ] Mobile responsive design
- [ ] Unit and integration tests

### 📄 License

MIT License

### 👥 Contributing

Issues and Pull Requests are welcome!

---

## 한국어

듣기 쓰기 연습을 통해 외국어 실력을 향상시키는 웹 애플리케이션입니다. 실생활 시나리오 대화를 통해 듣기 및 쓰기 과정에서 낯선 단어를 발견하고 기억할 수 있습니다.

### 🌍 지원 언어

- **🇨🇳 중국어** - 영어, 일본어, 한국어 학습
- **🇺🇸 영어** - 중국어, 일본어, 한국어 학습
- **🇰🇷 한국어** - 영어, 중국어, 일본어 학습
- **🇯🇵 일본어** - 영어, 중국어, 한국어 학습

### 🎯 핵심 기능

- **📚 시나리오 기반 학습**: 일상 생활, 직장, 여행 등 다양한 실생활 시나리오 제공
- **🎧 듣기 받아쓰기**: 오디오를 재생하고 들은 내용을 작성
- **⚡ 즉각 피드백**: 정답과 번역을 즉시 표시
- **📊 난이도 분류**: 쉬움, 보통, 어려움 세 가지 난이도로 분류된 문장
- **📈 진도 추적**: 사용자의 학습 진도와 성과 기록
- **🌐 다국어 인터페이스**: 중국어, 영어, 한국어, 일본어 인터페이스 지원

### 🛠️ 기술 스택

#### 백엔드
- **언어**: Go 1.21+
- **프레임워크**: Gin
- **데이터베이스**: MySQL + GORM
- **설정**: Viper (YAML)

#### 프론트엔드
- **프레임워크**: React 18 + TypeScript
- **UI 라이브러리**: Ant Design
- **스타일링**: Tailwind CSS
- **상태 관리**: Zustand
- **라우터**: React Router
- **HTTP**: Axios
- **오디오**: Howler.js

### 🚀 빠른 시작

#### 백엔드 서비스

```bash
cd backend
go mod download
go run cmd/main.go
```

백엔드 서비스는 `http://localhost:8080`에서 시작됩니다

#### 프론트엔드 애플리케이션

```bash
cd frontend
npm install
npm start
```

프론트엔드 애플리케이션은 `http://localhost:3000`에서 시작됩니다

### 📅 개발 로드맵

- [x] 프론트엔드 및 백엔드 프레임워크 설정
- [x] 기본 API 엔드포인트
- [x] 페이지 레이아웃 및 라우팅
- [x] 오디오 재생 기능
- [x] 데이터베이스 통합 (MySQL + GORM)
- [x] 엄격한 계층 아키텍처 구현
- [x] YAML 구성 관리
- [ ] 다국어 TTS 서비스 통합
- [ ] 사용자 인증 시스템
- [ ] 진도 통계 및 분석
- [ ] 다국어 인터페이스 지원
- [ ] 모바일 반응형 디자인
- [ ] 단위 및 통합 테스트

### 📄 라이선스

MIT License

### 👥 기여

이슈와 풀 리퀘스트를 환영합니다!

---

## 日本語

聞き取り練習を通じて外国語のスキルを向上させるWebアプリケーションです。実生活のシナリオ会話を通じて、聞き取りと書き取りの過程で不慣れな単語を発見し記憶することができます。

### 🌍 対応言語

- **🇨🇳 中国語** - 英語、日本語、韓国語を学習
- **🇺🇸 英語** - 中国語、日本語、韓国語を学習
- **🇰🇷 韓国語** - 英語、中国語、日本語を学習
- **🇯🇵 日本語** - 英語、中国語、韓国語を学習

### 🎯 主な機能

- **📚 シナリオベース学習**: 日常生活、職場、旅行など様々な実生活シナリオ
- **🎧 音声ディクテーション**: 音声を再生して聞いた内容を書き取り
- **⚡ 即時フィードバック**: 正解と翻訳を即座に表示
- **📊 難易度レベル**: 簡単、普通、難しいの3つの難易度に分類された文章
- **📈 進捗追跡**: ユーザーの学習進捗と成果を記録
- **🌐 多言語インターフェース**: 中国語、英語、韓国語、日本語のインターフェースをサポート

### 🛠️ 技術スタック

#### バックエンド
- **言語**: Go 1.21+
- **フレームワーク**: Gin
- **データベース**: MySQL + GORM
- **設定**: Viper (YAML)

#### フロントエンド
- **フレームワーク**: React 18 + TypeScript
- **UIライブラリ**: Ant Design
- **スタイリング**: Tailwind CSS
- **状態管理**: Zustand
- **ルーター**: React Router
- **HTTP**: Axios
- **オーディオ**: Howler.js

### 🚀 クイックスタート

#### バックエンドサービス

```bash
cd backend
go mod download
go run cmd/main.go
```

バックエンドサービスは `http://localhost:8080` で起動します

#### フロントエンドアプリケーション

```bash
cd frontend
npm install
npm start
```

フロントエンドアプリケーションは `http://localhost:3000` で起動します

### 📅 開発ロードマップ

- [x] フロントエンドとバックエンドのフレームワーク設定
- [x] 基本APIエンドポイント
- [x] ページレイアウトとルーティング
- [x] オーディオ再生機能
- [x] データベース統合 (MySQL + GORM)
- [x] 厳格なレイヤードアーキテクチャの実装
- [x] YAML設定管理
- [ ] 多言語TTSサービス統合
- [ ] ユーザー認証システム
- [ ] 進捗統計と分析
- [ ] 多言語インターフェースサポート
- [ ] モバイルレスポンシブデザイン
- [ ] ユニットテストと統合テスト

### 📄 ライセンス

MIT License

### 👥 貢献

IssueとPull Requestを歓迎します！

---

## 📞 Contact / 联系方式 / 연락처 / お問い合わせ

- GitHub: [guocong-bincai/voiceWriter](https://github.com/guocong-bincai/voiceWriter)
- Issues: [Report a bug](https://github.com/guocong-bincai/voiceWriter/issues)

---

**Made with ❤️ for language learners worldwide**
