package main

import (
	"log"
	"todo-app/internal/controllers"
	config "todo-app/internal/database"
	"todo-app/internal/repositories"
	"todo-app/internal/routes"
	"todo-app/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	router := gin.Default()
	err := godotenv.Load()
	config.InitFirestore()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	todoRepo := repositories.NewTodoRepository()
	todoService := services.NewTodoService(todoRepo)
	todoController := controllers.NewTodoController(todoService)
	routes.RegisterTodoRoutes(router, todoController)

	log.Println("Server running on http://localhost:9090")
	router.Run("localhost:9090")
}
