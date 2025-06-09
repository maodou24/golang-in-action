package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	e := gin.New()

	// 修正模板路径
	e.LoadHTMLFiles("index.html") // 改用 LoadHTMLFiles，直接加载 index.html
	e.Run(":8080")
}
