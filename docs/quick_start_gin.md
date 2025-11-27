# Gin 快速开始指南

## 1. 下载依赖

### 方式一：使用 go mod tidy（推荐）

```bash
# 在项目根目录执行
go mod tidy
```

这个命令会：
- 下载 `go.mod` 中声明的所有依赖
- 自动添加缺失的依赖
- 移除未使用的依赖
- 更新 `go.sum` 文件

### 方式二：使用 go get

```bash
# 下载 Gin 框架
go get github.com/gin-gonic/gin

# 或者下载所有依赖
go get ./...
```

### 方式三：使用 go mod download

```bash
# 下载所有依赖到本地缓存
go mod download
```

## 2. 验证依赖安装

```bash
# 查看已安装的依赖
go list -m all

# 查看 Gin 版本
go list -m github.com/gin-gonic/gin
```

## 3. 运行项目

```bash
# 方式一：直接运行
go run cmd/main.go

# 方式二：使用 Makefile
make run

# 方式三：构建后运行
go build -o bin/rich_go cmd/main.go
./bin/rich_go
```

## 4. 测试 API

启动服务器后，可以测试以下接口：

```bash
# 健康检查
curl http://localhost:8080/health

# 获取用户列表
curl http://localhost:8080/api/v1/users

# 获取单个用户
curl http://localhost:8080/api/v1/users/1

# 创建用户
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{"name":"张三","email":"zhangsan@example.com"}'

# 更新用户
curl -X PUT http://localhost:8080/api/v1/users/1 \
  -H "Content-Type: application/json" \
  -d '{"name":"李四","email":"lisi@example.com"}'

# 删除用户
curl -X DELETE http://localhost:8080/api/v1/users/1
```

## 5. 项目结构

```
rich_go/
├── cmd/
│   └── main.go              # 应用入口
├── internal/
│   ├── app/
│   │   └── app.go          # 应用主逻辑
│   └── server/
│       └── http_server.go  # HTTP 服务器（Gin）
├── go.mod                   # 依赖定义
└── go.sum                   # 依赖校验和
```

## 6. 常见问题

### 网络问题（国内）

如果下载依赖遇到网络问题，可以设置 Go 代理：

```bash
# 使用国内代理
go env -w GOPROXY=https://goproxy.cn,direct

# 或者使用多个代理
go env -w GOPROXY=https://goproxy.cn,https://proxy.golang.org,direct
```

### 依赖版本问题

如果遇到版本冲突：

```bash
# 查看依赖树
go mod graph

# 更新到最新版本
go get -u github.com/gin-gonic/gin

# 整理依赖
go mod tidy
```

## 下一步

- 📖 查看 [依赖管理文档](dependency_management.md) 了解企业级依赖管理
- 📖 查看 [Gin 官方文档](https://gin-gonic.com/docs/)
- 📖 查看 [框架选择指南](frameworks.md)

