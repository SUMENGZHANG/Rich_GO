# 项目设置指南

## ✅ 已完成的工作

1. ✅ 添加 Gin 依赖到 `go.mod`
2. ✅ 创建基于 Gin 的 HTTP 服务器代码
3. ✅ 更新应用主逻辑集成 HTTP 服务器
4. ✅ 创建依赖管理文档

## 📋 下一步操作

### 1. 下载依赖

由于网络问题，依赖已添加到 `go.mod`，但需要下载到本地：

```bash
# 方式一：使用 go mod tidy（推荐）
go mod tidy

# 方式二：如果网络有问题，设置代理后重试
go env -w GOPROXY=https://goproxy.cn,direct
go mod tidy
```

**说明**:
- `go mod tidy` 会自动下载 `go.mod` 中声明的所有依赖
- 会生成 `go.sum` 文件（依赖校验和，必须提交到版本控制）

### 2. 验证依赖安装

```bash
# 查看已安装的依赖
go list -m all

# 应该能看到 gin-gonic/gin
```

### 3. 运行项目

```bash
# 运行项目
go run cmd/main.go

# 或使用 Makefile
make run
```

服务器启动后，访问：
- http://localhost:8080/health - 健康检查
- http://localhost:8080/api/v1/users - 用户列表

## 📁 项目结构

```
rich_go/
├── cmd/
│   └── main.go                    # 应用入口
├── internal/
│   ├── app/
│   │   └── app.go                # 应用主逻辑
│   └── server/
│       └── http_server.go        # Gin HTTP 服务器 ⭐
├── go.mod                         # 依赖定义（已添加 Gin）
├── go.sum                         # 依赖校验和（运行 go mod tidy 后生成）
└── docs/
    ├── dependency_management.md   # 依赖管理指南
    └── quick_start_gin.md         # Gin 快速开始
```

## 🔍 Go 依赖管理机制

### Go Modules（go mod）

Go 使用 **Go Modules** 作为官方依赖管理工具：

1. **go.mod** - 定义模块和依赖
   - 模块名称
   - Go 版本要求
   - 直接依赖列表

2. **go.sum** - 依赖校验和
   - 记录依赖的哈希值
   - 确保依赖完整性
   - **必须提交到版本控制**

### 常用命令

```bash
# 添加依赖
go get github.com/gin-gonic/gin

# 下载依赖
go mod tidy          # 推荐：整理并下载
go mod download      # 只下载到缓存

# 查看依赖
go list -m all       # 查看所有依赖
go mod graph         # 查看依赖树

# 更新依赖
go get -u ./...      # 更新所有依赖
go get -u github.com/gin-gonic/gin  # 更新特定依赖

# 移除依赖
# 1. 从代码中删除 import
# 2. 运行 go mod tidy
```

## 🏢 企业级依赖管理最佳实践

### ✅ 推荐做法

1. **使用具体版本号**（生产环境）
   ```go
   require github.com/gin-gonic/gin v1.10.0
   ```

2. **提交 go.sum 文件**
   - go.sum 必须提交到版本控制
   - 确保团队使用相同版本的依赖

3. **定期更新依赖**
   ```bash
   go get -u ./...  # 更新所有依赖
   ```

4. **使用 go mod tidy**
   - 保持依赖整洁
   - 自动添加缺失的依赖
   - 移除未使用的依赖

5. **检查安全漏洞**
   ```bash
   go install golang.org/x/vuln/cmd/govulncheck@latest
   govulncheck ./...
   ```

### ❌ 避免做法

1. ❌ 在生产环境使用 `latest` 或 `master`
2. ❌ 忽略 go.sum 文件
3. ❌ 手动编辑 go.mod（除非必要）
4. ❌ 不更新过时的依赖

## 📚 相关文档

- [依赖管理详细指南](docs/dependency_management.md)
- [Gin 快速开始](docs/quick_start_gin.md)
- [框架选择指南](docs/frameworks.md)

## 🆘 常见问题

### 网络问题（国内）

```bash
# 设置 Go 代理
go env -w GOPROXY=https://goproxy.cn,direct
```

### 依赖下载失败

```bash
# 清理缓存后重试
go clean -modcache
go mod tidy
```

### 版本冲突

```bash
# 查看依赖树
go mod graph

# 更新到最新版本
go get -u github.com/gin-gonic/gin
```

## ✨ 下一步

1. 运行 `go mod tidy` 下载依赖
2. 运行 `go run cmd/main.go` 启动服务器
3. 测试 API 接口
4. 查看文档了解更多功能

