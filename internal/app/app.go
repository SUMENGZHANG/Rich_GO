package app

import (
	"fmt"
	"rich_go/internal/server"
)

// App 应用主结构
type App struct {
	Name       string
	Version    string
	HTTPServer *server.HTTPServer
}

// New 创建新的应用实例
func New() *App {
	return &App{
		Name:       "Rich_GO",
		Version:    "1.0.0",
		HTTPServer: server.NewHTTPServer("8080"),
	}
}

// Run 运行应用
func (a *App) Run() {
	fmt.Printf("应用启动中... [%s v%s]\n", a.Name, a.Version)
	fmt.Printf("HTTP 服务器启动在端口: %s\n", "8080")
	fmt.Println("访问 http://localhost:8080/health 查看健康状态")
	
	// 启动 HTTP 服务器
	if err := a.HTTPServer.Start(); err != nil {
		panic(fmt.Sprintf("启动 HTTP 服务器失败: %v", err))
	}
}

