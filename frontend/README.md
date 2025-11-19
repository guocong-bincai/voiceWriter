# VoiceWriter Frontend

基于 React + TypeScript + Ant Design 的英语听写练习前端应用

## 技术栈

- React 18
- TypeScript
- Ant Design (UI组件库)
- Tailwind CSS (样式)
- Zustand (状态管理)
- React Router (路由)
- Axios (HTTP客户端)
- Howler.js (音频播放)

## 项目结构

```
frontend/
├── public/              # 静态资源
├── src/
│   ├── components/     # 可复用组件
│   ├── pages/          # 页面组件
│   │   ├── HomePage.tsx       # 场景选择页
│   │   ├── ScenePage.tsx      # 句子列表页
│   │   └── PracticePage.tsx   # 听写练习页
│   ├── hooks/          # 自定义Hooks
│   │   └── useAudio.ts        # 音频播放Hook
│   ├── services/       # API服务
│   │   └── api.ts            # API接口定义
│   ├── store/          # 状态管理
│   │   └── index.ts          # Zustand Store
│   ├── types/          # TypeScript类型定义
│   │   └── index.ts
│   ├── assets/         # 资源文件
│   ├── App.tsx         # 根组件
│   └── index.tsx       # 入口文件
├── package.json
├── tsconfig.json
├── tailwind.config.js
└── postcss.config.js
```

## 快速开始

### 安装依赖

```bash
cd frontend
npm install
```

### 启动开发服务器

```bash
npm start
```

应用将在 `http://localhost:3000` 启动

### 构建生产版本

```bash
npm run build
```

## 环境变量

创建 `.env` 文件并配置：

```bash
REACT_APP_API_URL=http://localhost:8080/api/v1
```

## 功能页面

1. **首页 (/)** - 场景选择页，展示不同的学习场景（日常生活、工作职场、旅游出行等）
2. **场景页 (/scene/:id)** - 显示选定场景下的所有句子列表
3. **练习页 (/practice/:id)** - 听写练习界面，播放音频并输入答案

## 开发计划

- [x] 基础框架搭建
- [x] 页面路由配置
- [x] 状态管理
- [x] API接口封装
- [x] 音频播放功能
- [ ] 用户认证
- [ ] 进度追踪
- [ ] 统计分析
- [ ] 响应式优化
- [ ] 单元测试
