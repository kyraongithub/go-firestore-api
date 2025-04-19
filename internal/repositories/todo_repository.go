package repositories

import (
	"context"
	"errors"
	"log"
	config "todo-app/internal/database"
	"todo-app/internal/models"

	"cloud.google.com/go/firestore"
)

type TodoRepositoryInterface interface {
	GetTodos() ([]models.Todo, error)
	GetTodo(id string) (*models.Todo, error)
	AddTodo(todo models.Todo) (*models.Todo, error)
	ToggleTodoStatus(id string) (*models.Todo, error)
	DeleteTodo(id string) error
}

var ctx = context.Background()

type TodoRepository struct {
	Collection *firestore.CollectionRef
}

func NewTodoRepository() TodoRepositoryInterface {
	return &TodoRepository{
		Collection: config.FirestoreClient.Collection("todos"),
	}
}

func (t *TodoRepository) GetTodos() ([]models.Todo, error) {
	iter := t.Collection.Documents(ctx)
	var todos []models.Todo

	for {
		doc, err := iter.Next()
		if err != nil {
			break
		}
		var todo models.Todo
		if err := doc.DataTo(&todo); err != nil {
			log.Printf("Failed to convert document to todo: %v", err)
			continue
		}
		todo.ID = doc.Ref.ID
		todos = append(todos, todo)
	}

	return todos, nil
}

func (t *TodoRepository) GetTodo(id string) (*models.Todo, error) {
	doc, err := t.Collection.Doc(id).Get(ctx)
	if err != nil {
		return nil, errors.New("todo not found")
	}
	var todo models.Todo
	if err := doc.DataTo(&todo); err != nil {
		return nil, errors.New("failed to convert document to todo")
	}
	todo.ID = doc.Ref.ID
	return &todo, nil
}

func (t *TodoRepository) AddTodo(todo models.Todo) (*models.Todo, error) {
	newTodo := map[string]interface{}{
		"item":      todo.Item,
		"completed": todo.Completed,
	}
	ref, _, err := t.Collection.Add(ctx, newTodo)
	if err != nil {
		log.Println(err)
		return nil, errors.New("failed to add todo")
	}

	savedTodo := &models.Todo{
		ID:        ref.ID,
		Item:      todo.Item,
		Completed: todo.Completed,
	}

	return savedTodo, nil
}

func (t *TodoRepository) ToggleTodoStatus(id string) (*models.Todo, error) {
	todo, err := t.GetTodo(id)
	if err != nil {
		return nil, errors.New("todo not found")
	}

	_, err = t.Collection.Doc(id).Update(ctx, []firestore.Update{
		{Path: "completed", Value: !todo.Completed},
	})
	if err != nil {
		return nil, errors.New("failed to update todo")
	}

	todo.Completed = !todo.Completed
	return todo, nil
}

func (t *TodoRepository) DeleteTodo(id string) error {
	_, err := t.Collection.Doc(id).Delete(ctx)
	if err != nil {
		return errors.New("failed to delete todo")
	}
	return nil
}
