package app

import (
	"fmt"
	"log"
)

// App 应用主结构
type App struct {
	Name    string
	Version string
}

// New 创建新的应用实例
func New() *App {
	return &App{
		Name:    "Rich_GO",
		Version: "1.0.0",
	}
}

// Run 运行应用
func (a *App) Run() {
	fmt.Printf("应用启动中... [%s v%s]\n", a.Name, a.Version)
	// 在这里添加你的应用逻辑
	fmt.Println("应用运行中...")
}

