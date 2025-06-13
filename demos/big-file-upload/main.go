package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	e := gin.New()

	// 修正模板路径
	e.LoadHTMLFiles("index.html") // 改用 LoadHTMLFiles，直接加载 index.html

	e.Handle(http.MethodPost, "/upload/init", initUploadTask)
	e.Handle(http.MethodPost, "/upload/chunk", upload)
	e.Handle(http.MethodPost, "/upload/complete", complete)

	e.Run(":8080")
}
