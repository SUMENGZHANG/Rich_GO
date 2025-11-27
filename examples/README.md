# 框架使用示例

本目录包含常用框架的使用示例代码。

## REST API 示例

### Gin 示例
- `gin_example.go` - Gin 框架的基础使用示例

运行示例：
```bash
# 安装依赖
go get github.com/gin-gonic/gin

# 运行示例
go run examples/gin_example.go

# 测试接口
curl http://localhost:8080/health
curl http://localhost:8080/api/v1/users/1
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{"name":"张三","email":"zhangsan@example.com"}'
```

## gRPC 示例

### 前置准备

1. 安装 Protocol Buffers 编译器：
```bash
# macOS
brew install protobuf

# 安装 Go 插件
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

2. 生成代码：
```bash
# 在项目根目录执行
protoc --go_out=. --go-grpc_out=. examples/grpc_example.proto
```

### 运行示例

1. 启动服务器：
```bash
go run examples/grpc_server_example.go
```

2. 运行客户端：
```bash
go run examples/grpc_client_example.go
```

## 其他框架示例

更多框架示例请参考 [docs/frameworks.md](../docs/frameworks.md)

