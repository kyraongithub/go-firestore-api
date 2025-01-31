package controllers

import (
	"log"
	"net/http"
	"todo-app/internal/models"
	"todo-app/internal/services"

	"github.com/gin-gonic/gin"
)

type TodoController struct {
	Service services.TodoServiceInterface
}

func NewTodoController(service services.TodoServiceInterface) *TodoController {
	return &TodoController{Service: service}
}

func (tc *TodoController) GetTodos(c *gin.Context) {
	todos, err := tc.Service.GetTodos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "OK", "message": "get todos success", "code": "200", "data": todos})
}

func (tc *TodoController) GetTodo(c *gin.Context) {
	id := c.Param("id")
	todo, err := tc.Service.GetTodo(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "OK", "message": "get data success", "code": "200", "data": todo})
}

func (tc *TodoController) AddTodo(c *gin.Context) {
	var newTodo models.Todo
	if err := c.BindJSON(&newTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request payload"})
		return
	}

	todo, err := tc.Service.AddTodo(newTodo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "OK", "message": "data created succesfully", "code": "200", "data": todo})
}

func (tc *TodoController) ToggleTodoStatus(c *gin.Context) {
	id := c.Param("id")
	todo, err := tc.Service.ToggleTodoStatus(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "OK", "message": "data updated succesfully", "code": "200", "data": todo})
}

func (tc *TodoController) DeleteTodo(c *gin.Context) {
	id := c.Param("id")
	err := tc.Service.DeleteTodo(id)
	if err != nil {
		log.Fatalf("error")
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "OK", "message": "data deleted succesfully", "code": "200"})

}
