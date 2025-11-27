.PHONY: build run clean test fmt vet lint help

# 应用名称
APP_NAME=rich_go
BIN_DIR=bin
CMD_DIR=cmd

# 默认目标
help:
	@echo "可用的命令:"
	@echo "  make build    - 构建应用程序"
	@echo "  make run      - 运行应用程序"
	@echo "  make clean    - 清理构建文件"
	@echo "  make test     - 运行测试"
	@echo "  make fmt      - 格式化代码"
	@echo "  make vet      - 运行 go vet"
	@echo "  make lint     - 运行代码检查"
	@echo "  make deps     - 下载依赖"

# 构建应用
build:
	@echo "构建应用程序..."
	@mkdir -p $(BIN_DIR)
	@go build -o $(BIN_DIR)/$(APP_NAME) $(CMD_DIR)/main.go
	@echo "构建完成: $(BIN_DIR)/$(APP_NAME)"

# 运行应用
run:
	@echo "运行应用程序..."
	@go run $(CMD_DIR)/main.go

# 清理构建文件
clean:
	@echo "清理构建文件..."
	@rm -rf $(BIN_DIR)
	@go clean
	@echo "清理完成"

# 运行测试
test:
	@echo "运行测试..."
	@go test -v ./...

# 格式化代码
fmt:
	@echo "格式化代码..."
	@go fmt ./...

# 运行 go vet
vet:
	@echo "运行 go vet..."
	@go vet ./...

# 下载依赖
deps:
	@echo "下载依赖..."
	@go mod download
	@go mod tidy

# 代码检查（需要安装 golangci-lint）
lint:
	@echo "运行代码检查..."
	@if command -v golangci-lint > /dev/null; then \
		golangci-lint run; \
	else \
		echo "golangci-lint 未安装，跳过检查"; \
		echo "安装方法: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest"; \
	fi

