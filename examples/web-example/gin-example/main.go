package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
)

func main() {
	e := gin.New()

	e.Use(gin.Recovery())
	e.Use(gin.Logger())
	e.Use(cors.Default())
	e.Use(func(ctx *gin.Context) {
		ctx.Set("username", "admin")
		ctx.Set("userid", 1)
	})

	// use keys
	e.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, fmt.Sprintf("Hello %v", ctx.GetString("username")))
	})

	cal := e.Group("/calculate")
	// GET /path/sum?a=1&b=2
	cal.GET("/sum", func(ctx *gin.Context) {
		a := ctx.Query("a")
		b := ctx.Query("b")
		ctx.JSON(http.StatusOK, gin.H{"sum": sum(a, b)})
	})
	cal.POST("/sum", func(ctx *gin.Context) {
		type param struct {
			A int
			B int
		}
		var p param
		err := ctx.ShouldBindJSON(&p)
		if err != nil {
			err = ctx.Error(err)
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	})

	// GET /path/maodou
	e.GET("/hello/:name", func(c *gin.Context) {
		c.String(http.StatusOK, fmt.Sprintf("Hello %v", c.Param("name")))
	})

	e.POST("/login", func(ctx *gin.Context) {
		username := ctx.PostForm("username")
		password := ctx.PostForm("password")

		var result string
		if username == "maodou" && password == "maodou" {
			result = "login success"
		} else {
			result = "login fail, user or password error"
		}
		ctx.String(http.StatusOK, result)
	})

	err := e.Run(":80")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func sum(a, b string) int {
	i, _ := strconv.Atoi(a)
	j, _ := strconv.Atoi(b)
	return i + j
}
