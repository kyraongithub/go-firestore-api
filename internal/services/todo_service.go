package services

import (
	"todo-app/internal/models"
	"todo-app/internal/repositories"
)

type TodoServiceInterface interface {
	GetTodos() ([]models.Todo, error)
	GetTodo(id string) (*models.Todo, error)
	AddTodo(todo models.Todo) (*models.Todo, error)
	ToggleTodoStatus(id string) (*models.Todo, error)
	DeleteTodo(id string) error
}

type TodoService struct {
	Repo repositories.TodoRepositoryInterface
}

func NewTodoService(repo repositories.TodoRepositoryInterface) *TodoService {
	return &TodoService{Repo: repo}
}

func (s *TodoService) GetTodos() ([]models.Todo, error) {
	return s.Repo.GetTodos()
}

func (s *TodoService) GetTodo(id string) (*models.Todo, error) {
	return s.Repo.GetTodo(id)
}

func (s *TodoService) AddTodo(todo models.Todo) (*models.Todo, error) {
	return s.Repo.AddTodo(todo)
}

func (s *TodoService) ToggleTodoStatus(id string) (*models.Todo, error) {
	return s.Repo.ToggleTodoStatus(id)
}

func (s *TodoService) DeleteTodo(id string) error {
	err := s.Repo.DeleteTodo(id)
	if err != nil {
		return err
	}
	return nil
}
