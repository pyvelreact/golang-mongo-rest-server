package todoroute

import (
	"ccgolang/server/api/todo/controller"
	"github.com/labstack/echo"
)

func Init(e *echo.Echo) {
	e.GET("/api/todos", todocontroller.GetAll)
	e.GET("/api/todos/:id", todocontroller.GetById)
	e.POST("/api/todos", todocontroller.NewTodo)
	e.DELETE("/api/todos/:id", todocontroller.RemoveTodo)
}
