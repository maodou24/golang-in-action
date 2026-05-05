package controller

import (
	"mvc-sample/repo"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserController(ctx *gin.Context) {
	name, ok := ctx.GetQuery("name")
	// 前后端分离项目，controller决定视图逻辑
	// 例如 如果有name参数，则查询单个用户，否则查询所有用户
	if ok {
		user, err := repo.GetUserRepo().FindByName(name)
		if err != nil {
			ctx.JSON(http.StatusNotFound, "User not found")
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"data": user})
		return
	}

	users, err := repo.GetUserRepo().FindAll()
	if err != nil {
		ctx.JSON(http.StatusNotFound, "User not found")
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": users})
}
