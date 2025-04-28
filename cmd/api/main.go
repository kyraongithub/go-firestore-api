package main

import (
	"fmt"
	"log"
	"os"
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
	port := os.Getenv("PROJECT_PORT")
	if port == "" {
		port = "9090"
	}
	log.Println(port)

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	todoRepo := repositories.NewTodoRepository()
	todoService := services.NewTodoService(todoRepo)
	todoController := controllers.NewTodoController(todoService)
	routes.RegisterTodoRoutes(router, todoController)

	log.Printf("Server running on http://localhost:%s", port)
	router.Run(fmt.Sprintf(":%s", port))
}
