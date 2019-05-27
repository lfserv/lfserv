package main

import (
	"github.com/gramework/gramework"
)

func main() {

	app := gramework.New()
	app.GET("/", "hello")
	app.GET("/:name", func(ctx *gramework.Context) {
		name := ctx.RouteArg("name")
		ctx.WriteString("hello, " + name)
	})
	app.ListenAndServe(":8080")
}
