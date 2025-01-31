package routes

import (
	"todo-app/internal/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterTodoRoutes(router *gin.Engine, controller *controllers.TodoController) {
	todoRoutes := router.Group("/todos")
	{
		todoRoutes.GET("/", controller.GetTodos)
		todoRoutes.GET("/:id", controller.GetTodo)
		todoRoutes.POST("/", controller.AddTodo)
		todoRoutes.PATCH("/:id", controller.ToggleTodoStatus)
		todoRoutes.DELETE("/:id", controller.DeleteTodo)
	}
}
