# Go 企业级项目依赖管理指南

## Go 依赖管理机制

Go 使用 **Go Modules**（go mod）作为官方的依赖管理工具，从 Go 1.11 版本开始引入，Go 1.13 后成为默认依赖管理方式。

## 核心概念

### 1. go.mod 文件

`go.mod` 是 Go 模块的配置文件，定义了：
- 模块名称（module path）
- Go 版本要求
- 直接依赖和版本

```go
module rich_go

go 1.21

require (
    github.com/gin-gonic/gin v1.10.0
)
```

### 2. go.sum 文件

`go.sum` 文件记录依赖的校验和，用于：
- 验证依赖的完整性
- 确保依赖版本一致性
- **应该提交到版本控制系统**

## 常用命令

### 添加依赖

```bash
# 方式一：直接添加（推荐）
go get github.com/gin-gonic/gin

# 方式二：指定版本
go get github.com/gin-gonic/gin@v1.10.0

# 方式三：指定最新版本
go get github.com/gin-gonic/gin@latest

# 方式四：更新到最新版本
go get -u github.com/gin-gonic/gin
```

### 下载依赖

```bash
# 下载所有依赖
go mod download

# 整理依赖（添加缺失的，移除未使用的）
go mod tidy

# 验证依赖
go mod verify
```

### 查看依赖

```bash
# 查看所有依赖
go list -m all

# 查看特定依赖
go list -m github.com/gin-gonic/gin

# 查看依赖树
go mod graph
```

### 移除依赖

```bash
# 从代码中移除导入后，运行
go mod tidy

# 或者手动移除 go.mod 中的依赖，然后运行
go mod tidy
```

## 企业级项目依赖管理最佳实践

### 1. 版本管理策略

#### 语义化版本（Semantic Versioning）

Go Modules 使用语义化版本：
- `v1.2.3` - 主版本.次版本.修订版本
- `v0.1.0` - 开发版本
- `v1.2.3-beta.1` - 预发布版本

#### 版本选择建议

```go
// ✅ 推荐：使用具体版本（生产环境）
require github.com/gin-gonic/gin v1.10.0

// ✅ 推荐：使用最新补丁版本
require github.com/gin-gonic/gin v1.10

// ⚠️ 谨慎：使用最新版本（开发环境）
require github.com/gin-gonic/gin latest

// ❌ 不推荐：使用 master/main（不稳定）
require github.com/gin-gonic/gin master
```

### 2. 依赖分类管理

#### 直接依赖 vs 间接依赖

- **直接依赖**：项目代码中直接导入的包
- **间接依赖**：直接依赖的依赖

```bash
# 查看直接依赖
go list -m -f '{{if not .Indirect}}{{.}}{{end}}' all

# 查看间接依赖
go list -m -f '{{if .Indirect}}{{.}}{{end}}' all
```

### 3. 依赖更新策略

#### 定期更新依赖

```bash
# 更新所有依赖到最新版本
go get -u ./...

# 更新特定依赖
go get -u github.com/gin-gonic/gin

# 更新到最新补丁版本（不更新主版本）
go get -u=patch ./...
```

#### 更新检查清单

1. ✅ 查看更新日志（CHANGELOG）
2. ✅ 运行测试确保兼容性
3. ✅ 检查是否有破坏性变更
4. ✅ 更新文档和配置

### 4. 依赖安全

#### 检查安全漏洞

```bash
# 使用 govulncheck（Go 官方工具）
go install golang.org/x/vuln/cmd/govulncheck@latest
govulncheck ./...

# 使用第三方工具
# - Snyk
# - Dependabot (GitHub)
# - OWASP Dependency-Check
```

### 5. 私有依赖管理

#### 使用私有 Go Module 代理

```bash
# 设置 GOPRIVATE（不通过公共代理）
go env -w GOPRIVATE=github.com/yourcompany/*

# 配置 Git 认证
git config --global url."https://token@github.com/yourcompany".insteadOf "https://github.com/yourcompany"
```

#### 使用 Go Module Proxy

```bash
# 设置私有代理
go env -w GOPROXY=https://proxy.golang.org,direct
go env -w GOSUMDB=sum.golang.org

# 企业内网代理
go env -w GOPROXY=https://goproxy.company.com,direct
```

### 6. 依赖锁定

#### go.sum 文件

- ✅ **必须提交到版本控制**
- ✅ 确保团队使用相同版本的依赖
- ✅ 防止依赖被篡改

```bash
# 验证 go.sum
go mod verify
```

### 7. 依赖最小化原则

#### 只添加必要的依赖

```bash
# 定期清理未使用的依赖
go mod tidy

# 检查未使用的依赖
go mod why -m <module>
```

## 项目结构示例

### 标准 Go 项目依赖管理

```
rich_go/
├── go.mod              # 模块定义和直接依赖
├── go.sum              # 依赖校验和（必须提交）
├── cmd/
│   └── main.go        # 主程序入口
├── internal/          # 内部代码（不导出）
│   ├── app/
│   ├── server/
│   └── handler/
├── pkg/               # 可导出包
└── vendor/            # 可选：依赖副本（用于离线构建）
```

### vendor 目录（可选）

```bash
# 创建 vendor 目录（将依赖复制到项目）
go mod vendor

# 使用 vendor 构建
go build -mod=vendor ./cmd/main.go
```

**使用场景**:
- 离线构建
- 确保构建一致性
- CI/CD 环境

## 常见问题解决

### 1. 网络问题

```bash
# 设置 Go 代理（国内推荐）
go env -w GOPROXY=https://goproxy.cn,direct

# 或者使用多个代理
go env -w GOPROXY=https://goproxy.cn,https://proxy.golang.org,direct
```

### 2. 版本冲突

```bash
# 查看依赖冲突
go mod why -m <module>

# 使用 replace 指令解决冲突
go mod edit -replace github.com/old/module=github.com/new/module@v1.0.0
```

### 3. 依赖更新失败

```bash
# 清理模块缓存
go clean -modcache

# 重新下载
go mod download
```

### 4. 私有仓库认证

```bash
# 配置 Git 凭据
git config --global credential.helper store

# 或使用 .netrc 文件
machine github.com
login your-username
password your-token
```

## 企业级最佳实践总结

### ✅ 推荐做法

1. **使用具体版本号**（生产环境）
2. **定期更新依赖**（每月或每季度）
3. **提交 go.sum 文件**到版本控制
4. **使用 go mod tidy** 保持依赖整洁
5. **检查安全漏洞**（使用 govulncheck）
6. **文档化依赖选择理由**
7. **使用依赖锁定**（go.sum）

### ❌ 避免做法

1. ❌ 使用 `latest` 或 `master`（生产环境）
2. ❌ 忽略 go.sum 文件
3. ❌ 手动编辑 go.mod（除非必要）
4. ❌ 忽略安全漏洞
5. ❌ 添加不必要的依赖
6. ❌ 不更新过时的依赖

## 依赖管理工具

### 官方工具

- `go mod` - 官方依赖管理
- `govulncheck` - 安全漏洞检查

### 第三方工具

- **golangci-lint** - 代码检查（包括依赖）
- **modd** - 文件监控和自动构建
- **gopls** - Go 语言服务器（IDE 支持）

## 示例：添加 Gin 依赖

```bash
# 1. 添加依赖
go get github.com/gin-gonic/gin

# 2. 整理依赖
go mod tidy

# 3. 验证依赖
go mod verify

# 4. 查看依赖
go list -m all
```

## 参考资源

- [Go Modules 官方文档](https://go.dev/ref/mod)
- [Go Modules Wiki](https://github.com/golang/go/wiki/Modules)
- [Go 依赖管理最佳实践](https://go.dev/doc/modules/managing-dependencies)
- [Go 安全漏洞检查](https://go.dev/blog/vuln)

