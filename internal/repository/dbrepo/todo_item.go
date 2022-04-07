package dbrepo

import (
	"context"

	"github.com/Shambou/todolist/internal/models"
)

// GetTodoItems returns all todo_items based on completed bool
func (d *PostgresDBRepo) GetTodoItems(ctx context.Context, completed bool) ([]models.TodoItem, error) {
	var items []models.TodoItem

	result := d.DB.Where("completed = ?", completed).Find(&items)

	if result.Error != nil {
		return items, d.DB.Error
	}

	return items, nil
}

func (d *PostgresDBRepo) InsertItem(ctx context.Context, item models.TodoItem) (models.TodoItem, error) {
	result := d.DB.Create(&item)
	if result.Error != nil {
		return item, d.DB.Error
	}
	d.DB.Last(&item)

	return item, nil
}
