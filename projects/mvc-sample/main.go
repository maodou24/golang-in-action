package main

import (
	"fmt"
	"mvc-sample/controller"
	"mvc-sample/db"
	"mvc-sample/repo"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	database := db.GetDb()

	repo.InitUserRepo(database)

	e := gin.New()

	// middleware support
	e.Use(gin.Recovery())
	e.Use(gin.Logger())
	e.Use(cors.Default())

	e.GET("/user", controller.UserController)

	err := e.Run(":8080")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
