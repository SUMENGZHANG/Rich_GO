# 项目架构详解

## 一、internal 文件夹的作用

### 1.1 internal 的设计理念

`internal` 是 Go 语言的特殊目录，具有**包可见性限制**的特性：

- ✅ **私有包**: `internal` 目录下的包**只能被父目录及其子目录导入**
- ✅ **外部不可见**: 其他项目无法导入 `internal` 下的包
- ✅ **封装内部实现**: 隐藏实现细节，只暴露必要的接口

### 1.2 internal 的访问规则

```
rich_go/
├── cmd/
│   └── main.go          ✅ 可以导入 internal/*
├── internal/
│   ├── app/             ✅ 可以导入 internal/*
│   └── server/          ✅ 可以导入 internal/*
├── pkg/                 ✅ 可以导入 internal/*
└── 外部项目              ❌ 无法导入 internal/*
```

**示例**:
```go
// ✅ 可以 - 在项目内部
// cmd/main.go
import "rich_go/internal/app"

// ❌ 不可以 - 外部项目
// 其他项目的代码
import "rich_go/internal/app"  // 编译错误！
```

### 1.3 为什么使用 internal？

1. **封装内部实现**
   - 隐藏实现细节
   - 只暴露必要的公共接口
   - 降低耦合度

2. **防止外部依赖**
   - 外部项目无法导入内部包
   - 避免破坏性变更影响外部用户
   - 保持 API 稳定性

3. **清晰的代码组织**
   - 区分公共 API 和内部实现
   - 提高代码可维护性

## 二、项目架构总览

### 2.1 目录结构

```
rich_go/
├── cmd/                    # 应用程序入口
│   └── main.go            # 主程序入口点
│
├── internal/              # 私有应用代码（外部不可导入）
│   ├── app/              # 应用核心逻辑
│   │   └── app.go        # 应用主结构
│   └── server/           # HTTP 服务器
│       └── http_server.go # Gin HTTP 服务器实现
│
├── pkg/                   # 可被外部使用的库代码
│                         # （目前为空，可放置公共库）
│
├── api/                   # API 定义文件
│                         # （Proto 文件、OpenAPI 定义等）
│
├── configs/              # 配置文件
│   └── config.example.yaml # 配置示例
│
├── docs/                 # 项目文档
│   ├── frameworks.md
│   ├── dependency_management.md
│   └── ...
│
├── examples/             # 示例代码
│   ├── gin_example.go
│   └── grpc_example.proto
│
├── scripts/              # 脚本文件
│
├── go.mod                # Go 模块定义
├── go.sum                # 依赖校验和
├── Makefile              # 构建脚本
└── README.md             # 项目说明
```

### 2.2 架构层次图

```
┌─────────────────────────────────────────────────────────┐
│                    应用层 (Application Layer)            │
│  ┌───────────────────────────────────────────────────┐  │
│  │  cmd/main.go                                      │  │
│  │  - 程序入口点                                     │  │
│  │  - 初始化应用                                     │  │
│  └───────────────────────────────────────────────────┘  │
└────────────────────┬────────────────────────────────────┘
                     │
                     ↓
┌─────────────────────────────────────────────────────────┐
│                   业务逻辑层 (Business Logic Layer)       │
│  ┌───────────────────────────────────────────────────┐  │
│  │  internal/app/                                     │  │
│  │  - App 结构体（应用主结构）                        │  │
│  │  - 应用生命周期管理                                │  │
│  │  - 协调各个组件                                    │  │
│  └───────────────────────────────────────────────────┘  │
└────────────────────┬────────────────────────────────────┘
                     │
                     ↓
┌─────────────────────────────────────────────────────────┐
│                   服务层 (Service Layer)                  │
│  ┌───────────────────────────────────────────────────┐  │
│  │  internal/server/                                  │  │
│  │  - HTTP 服务器实现                                  │  │
│  │  - 路由注册                                        │  │
│  │  - 中间件配置                                      │  │
│  │  - 请求处理函数                                    │  │
│  └───────────────────────────────────────────────────┘  │
└────────────────────┬────────────────────────────────────┘
                     │
                     ↓
┌─────────────────────────────────────────────────────────┐
│                   框架层 (Framework Layer)               │
│  ┌───────────────────────────────────────────────────┐  │
│  │  github.com/gin-gonic/gin                          │  │
│  │  - HTTP 框架                                       │  │
│  │  - 路由引擎                                        │  │
│  │  - 中间件系统                                      │  │
│  └───────────────────────────────────────────────────┘  │
└────────────────────┬────────────────────────────────────┘
                     │
                     ↓
┌─────────────────────────────────────────────────────────┐
│                   运行时层 (Runtime Layer)                │
│  ┌───────────────────────────────────────────────────┐  │
│  │  Go 标准库 + Go 运行时                             │  │
│  │  - net/http                                        │  │
│  │  - Goroutine 调度器                                │  │
│  │  - 网络轮询器                                      │  │
│  └───────────────────────────────────────────────────┘  │
└─────────────────────────────────────────────────────────┘
```

## 三、各层详细说明

### 3.1 cmd/ - 应用入口层

**职责**: 程序入口点，初始化应用

```go
// cmd/main.go
package main

import "rich_go/internal/app"

func main() {
    // 1. 创建应用实例
    application := app.New()
    
    // 2. 运行应用
    application.Run()
}
```

**特点**:
- ✅ 最小化代码，只负责启动
- ✅ 不包含业务逻辑
- ✅ 可以包含多个入口（如 `cmd/server/main.go`, `cmd/cli/main.go`）

### 3.2 internal/app/ - 应用核心层

**职责**: 应用主结构，协调各个组件

```go
// internal/app/app.go
package app

type App struct {
    Name       string
    Version    string
    HTTPServer *server.HTTPServer  // 组合 HTTP 服务器
}

func New() *App {
    return &App{
        Name:       "Rich_GO",
        Version:    "1.0.0",
        HTTPServer: server.NewHTTPServer("8080"),
    }
}

func (a *App) Run() {
    // 启动各个组件
    a.HTTPServer.Start()
}
```

**特点**:
- ✅ 应用生命周期管理
- ✅ 组件协调和初始化
- ✅ 可以扩展其他组件（数据库、缓存等）

**未来可以扩展**:
```go
type App struct {
    Name       string
    Version    string
    HTTPServer *server.HTTPServer
    Database   *database.DB        // 数据库
    Cache      *cache.Cache        // 缓存
    Config     *config.Config      // 配置
}
```

### 3.3 internal/server/ - HTTP 服务器层

**职责**: HTTP 服务器实现，路由和请求处理

```go
// internal/server/http_server.go
package server

type HTTPServer struct {
    router *gin.Engine
    port   string
}

func NewHTTPServer(port string) *HTTPServer {
    router := gin.Default()
    setupMiddleware(router)  // 设置中间件
    setupRoutes(router)       // 注册路由
    return &HTTPServer{router: router, port: port}
}

func (s *HTTPServer) Start() error {
    return s.router.Run(":" + s.port)
}
```

**特点**:
- ✅ HTTP 服务器封装
- ✅ 路由注册和管理
- ✅ 中间件配置
- ✅ 请求处理函数

**当前路由结构**:
```
GET  /health              → healthCheck
GET  /api/v1/users        → listUsers
GET  /api/v1/users/:id    → getUser
POST /api/v1/users        → createUser
PUT  /api/v1/users/:id    → updateUser
DELETE /api/v1/users/:id  → deleteUser
```

### 3.4 pkg/ - 公共库层

**职责**: 可被外部项目使用的公共库

**当前状态**: 空目录

**未来可以放置**:
- 公共工具函数
- 可复用的业务逻辑
- 第三方集成库

**示例**:
```
pkg/
├── utils/          # 工具函数
│   └── string.go
├── errors/         # 错误定义
│   └── errors.go
└── validator/      # 验证器
    └── validator.go
```

### 3.5 api/ - API 定义层

**职责**: API 接口定义文件

**当前状态**: 空目录

**未来可以放置**:
- Protocol Buffers 定义（`.proto`）
- OpenAPI/Swagger 定义（`.yaml`）
- GraphQL Schema（`.graphql`）

**示例**:
```
api/
├── proto/          # gRPC 定义
│   └── user.proto
└── openapi/        # REST API 定义
    └── api.yaml
```

### 3.6 configs/ - 配置层

**职责**: 配置文件

```
configs/
└── config.example.yaml  # 配置示例
```

**特点**:
- ✅ 提供配置示例
- ✅ 实际配置不提交到版本控制（.gitignore）

## 四、代码调用关系

### 4.1 调用链

```
main()
  │
  ├─> app.New()
  │     │
  │     └─> server.NewHTTPServer()
  │           ├─> gin.Default()
  │           ├─> setupMiddleware()
  │           └─> setupRoutes()
  │
  └─> app.Run()
        │
        └─> server.Start()
              │
              └─> router.Run()
                    │
                    └─> http.ListenAndServe()
```

### 4.2 依赖关系图

```
┌─────────┐
│  main   │
└────┬────┘
     │ import
     ↓
┌─────────────┐
│ internal/app│
└────┬────────┘
     │ import
     ↓
┌──────────────┐
│internal/server│
└──────┬───────┘
       │ import
       ↓
┌──────────────┐
│ gin-gonic/gin│
└──────────────┘
```

### 4.3 包导入规则

```go
// ✅ cmd/main.go 可以导入
import "rich_go/internal/app"

// ✅ internal/app/app.go 可以导入
import "rich_go/internal/server"

// ✅ internal/server/http_server.go 可以导入
import "github.com/gin-gonic/gin"

// ❌ 外部项目无法导入 internal/*
import "rich_go/internal/app"  // 编译错误！
```

## 五、架构设计原则

### 5.1 分层架构

- **清晰的分层**: 每层职责明确
- **依赖方向**: 上层依赖下层，下层不依赖上层
- **接口隔离**: 通过接口定义交互

### 5.2 封装原则

- **internal 封装**: 内部实现隐藏在 `internal/` 下
- **最小暴露**: 只暴露必要的公共接口
- **接口抽象**: 通过接口定义依赖关系

### 5.3 扩展性

- **组件化**: 各个组件独立，易于扩展
- **可替换**: 可以替换实现而不影响其他部分
- **可测试**: 各层可以独立测试

## 六、未来扩展建议

### 6.1 建议的目录结构扩展

```
rich_go/
├── cmd/
│   ├── server/main.go      # HTTP 服务器入口
│   └── cli/main.go         # CLI 工具入口
│
├── internal/
│   ├── app/                # 应用核心
│   ├── server/             # HTTP 服务器
│   ├── handler/            # 请求处理函数
│   │   └── user_handler.go
│   ├── service/            # 业务逻辑层
│   │   └── user_service.go
│   ├── repository/         # 数据访问层
│   │   └── user_repo.go
│   ├── model/              # 数据模型
│   │   └── user.go
│   ├── middleware/         # 自定义中间件
│   │   └── auth.go
│   └── config/             # 配置管理
│       └── config.go
│
├── pkg/
│   ├── utils/              # 工具函数
│   ├── errors/             # 错误定义
│   └── validator/          # 验证器
│
├── api/
│   ├── proto/              # gRPC 定义
│   └── openapi/            # REST API 定义
│
└── configs/
    └── config.yaml         # 实际配置（不提交）
```

### 6.2 分层架构扩展

```
┌─────────────────────────────────────────┐
│  Handler Layer (internal/handler/)      │  处理 HTTP 请求
├─────────────────────────────────────────┤
│  Service Layer (internal/service/)      │  业务逻辑
├─────────────────────────────────────────┤
│  Repository Layer (internal/repository/)│  数据访问
├─────────────────────────────────────────┤
│  Model Layer (internal/model/)         │  数据模型
└─────────────────────────────────────────┘
```

## 七、最佳实践总结

### 7.1 internal 使用原则

1. ✅ **私有实现放在 internal/**
   - 内部实现细节
   - 不希望外部依赖的代码

2. ✅ **公共 API 放在 pkg/**
   - 可被外部项目使用
   - 需要保持 API 稳定性

3. ✅ **入口放在 cmd/**
   - 程序入口点
   - 最小化代码

### 7.2 项目组织原则

1. ✅ **按功能分层**: Handler → Service → Repository
2. ✅ **按模块分组**: user/, order/, product/
3. ✅ **依赖注入**: 通过接口定义依赖
4. ✅ **配置外部化**: 配置文件独立管理

## 八、当前项目架构总结

### 8.1 当前架构特点

- ✅ **简洁清晰**: 三层架构（cmd → app → server）
- ✅ **职责明确**: 每层职责单一
- ✅ **易于扩展**: 可以轻松添加新组件
- ✅ **符合规范**: 遵循 Go 项目布局标准

### 8.2 当前代码流程

```
main() 
  → 创建 App
    → 创建 HTTPServer
      → 初始化 Gin
      → 注册路由
  → 启动服务器
    → 监听端口
    → 处理请求
```

### 8.3 internal 的作用总结

1. **封装内部实现**: 隐藏实现细节
2. **防止外部依赖**: 外部项目无法导入
3. **清晰的代码组织**: 区分公共和私有代码
4. **保持 API 稳定**: 内部变更不影响外部用户

## 总结

**internal 文件夹的作用**:
- 🔒 **私有包**: 只能被项目内部导入
- 🎯 **封装实现**: 隐藏内部实现细节
- 🛡️ **保护 API**: 防止外部依赖内部实现
- 📦 **代码组织**: 清晰区分公共和私有代码

**项目架构**:
- 📁 **cmd/**: 程序入口
- 📁 **internal/**: 私有实现（app、server）
- 📁 **pkg/**: 公共库（未来扩展）
- 📁 **api/**: API 定义（未来扩展）
- 📁 **configs/**: 配置文件

这是一个符合 Go 最佳实践的项目架构！

