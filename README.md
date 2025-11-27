# Rich_GO

一个使用 Go 语言开发的项目。

## 项目结构

```
Rich_GO/
├── cmd/              # 应用程序入口
│   └── main.go       # 主程序入口
├── internal/         # 私有应用代码
│   └── app/          # 应用核心逻辑
├── pkg/              # 可以被外部应用使用的库代码
├── api/              # API 定义文件
├── configs/          # 配置文件
├── scripts/          # 脚本文件
├── docs/             # 项目文档
│   └── frameworks.md # 框架选择指南
├── examples/         # 示例代码
│   ├── gin_example.go
│   ├── grpc_example.proto
│   └── README.md
├── go.mod            # Go 模块定义文件
└── README.md         # 项目说明文档
```

## 环境要求

- Go 1.21 或更高版本

## 安装 Go

### macOS
```bash
# 使用 Homebrew 安装
brew install go

# 或者从官网下载安装包
# https://golang.org/dl/
```

### Linux
```bash
# Ubuntu/Debian
sudo apt-get update
sudo apt-get install golang-go

# 或者从官网下载安装包
```

### Windows
从 [Go 官网](https://golang.org/dl/) 下载 Windows 安装包并安装。

## 快速开始

### 1. 安装依赖

```bash
# 下载所有依赖（推荐）
go mod tidy

# 或者只下载依赖到缓存
go mod download
```

### 2. 运行项目

```bash
# 方式一：直接运行
go run cmd/main.go

# 方式二：使用 Makefile
make run

# 方式三：构建后运行
make build
./bin/rich_go
```

### 3. 测试 API

服务器启动后（默认端口 8080），可以测试：

```bash
# 健康检查
curl http://localhost:8080/health

# 获取用户列表
curl http://localhost:8080/api/v1/users
```

📖 **详细使用指南**: 请查看 [docs/quick_start_gin.md](docs/quick_start_gin.md)

## 开发指南

### 项目结构

- `cmd/` - 应用程序入口点
- `internal/` - 私有应用代码，不会被外部导入
  - `app/` - 应用核心逻辑
  - `server/` - HTTP 服务器（基于 Gin）
- `pkg/` - 可以被外部应用使用的库代码
- `api/` - API 接口定义
- `configs/` - 配置文件目录

### 依赖管理

本项目使用 **Go Modules** 进行依赖管理：

- `go.mod` - 模块定义和依赖声明
- `go.sum` - 依赖校验和（必须提交到版本控制）

**常用命令**:
```bash
# 添加依赖
go get github.com/package/name

# 下载依赖
go mod tidy

# 查看依赖
go list -m all
```

📖 **详细依赖管理指南**: 请查看 [docs/dependency_management.md](docs/dependency_management.md)

🔍 **Go 运行时架构与程序执行流程**: 请查看 [docs/go_runtime_architecture.md](docs/go_runtime_architecture.md)

📊 **程序执行流程图解**: 请查看 [docs/program_execution_flow.md](docs/program_execution_flow.md)

## 框架选择

企业级 Go 服务端开发常用的框架选择：

### REST API 框架
- **Gin** - 最受欢迎，性能优秀，生态丰富 ⭐⭐⭐⭐⭐
- **Echo** - 高性能，API 设计优雅 ⭐⭐⭐⭐
- **Hertz** - 字节跳动开源，极致性能（基于 Netpoll） ⭐⭐⭐⭐⭐
- **Fiber** - 极致性能，类似 Express.js ⭐⭐⭐⭐
- **Chi** - 轻量级，标准库兼容 ⭐⭐⭐⭐

### RPC 框架
- **gRPC** - Google 官方，行业标准 ⭐⭐⭐⭐⭐
- **Kitex** - 字节跳动开源，高性能微服务框架 ⭐⭐⭐⭐⭐
- **RPCX** - 功能丰富的 RPC 框架 ⭐⭐⭐⭐
- **TarsGo** - 腾讯开源，完整服务治理 ⭐⭐⭐⭐

📖 **详细框架选择指南**: 请查看 [docs/frameworks.md](docs/frameworks.md)

📊 **Gin vs Echo 详细对比**: 请查看 [docs/gin_vs_echo.md](docs/gin_vs_echo.md)

🏢 **字节跳动内部框架选择**: 请查看 [docs/byte_dance_frameworks.md](docs/byte_dance_frameworks.md)

### 推荐组合
- **方案一（最常用）**: Gin + gRPC
- **方案二（高性能）**: Echo + gRPC
- **方案三（微服务）**: Kitex + Hertz（字节跳动方案）

## 许可证

[在此添加许可证信息]

