package main

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"os"
)

func main() {
	app := iris.New()
	app.Use(iris.Compression)
	app.Use(logger.New())

	//i18n.New().Load()

	app.Get("/", func(ctx iris.Context) {
		ctx.HTML("Hello <strong>%s</strong>!", "World")
	})

	cal := app.Party("/calculate")
	{
		cal.Get("/sum", func(ctx iris.Context) {
			a, _ := ctx.URLParamInt("a")
			b, _ := ctx.URLParamInt("b")
			ctx.JSON(iris.Map{"sum": a + b})
		})
		cal.Post("/sum", func(ctx iris.Context) {
			type param struct {
				A int `json:"a"`
				B int `json:"b"`
			}
			var p param
			err := ctx.ReadJSON(&p)
			if err != nil {
				ctx.JSON(iris.Map{"error": err.Error()})
				return
			}
			ctx.JSON(iris.Map{"sum": p.A + p.B})
		})
	}

	err := app.Listen(":80")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
