package main

import (
	"fmt"
	"rich_go/internal/app"
)

func main() {
	fmt.Println("欢迎使用 Rich_GO 项目!")
	
	// 初始化应用
	application := app.New()
	application.Run()
}

