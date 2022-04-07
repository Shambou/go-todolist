package repository

import (
	"context"

	"github.com/Shambou/todolist/internal/models"
)

type DatabaseRepo interface {
	GetTodoItems(ctx context.Context, id bool) ([]models.TodoItem, error)
	InsertItem(ctx context.Context, item models.TodoItem) (models.TodoItem, error)
}
