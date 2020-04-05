package main

import (
	"github.com/labstack/echo"
	"github.com/todo20200404/handlers"
)

func main() {
	e := echo.New()
	e.File("/", "public/index.html")
	e.GET("tasks", handlers.GetTasks)

	e.Start(":8888")
}
