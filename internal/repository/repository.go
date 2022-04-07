package repository

import (
	"context"

	"github.com/Shambou/todolist/internal/models"
)

type DatabaseRepo interface {
	GetItems(ctx context.Context, completed bool) ([]models.TodoItem, error)
	InsertItem(ctx context.Context, item models.TodoItem) (models.TodoItem, error)
	GetItemById(ctx context.Context, id int) (models.TodoItem, error)
	UpdateItem(ctx context.Context, existingItem models.TodoItem, data map[string]interface{}) (models.TodoItem, error)
}
