package dbrepo

import (
	"context"

	"github.com/Shambou/todolist/internal/models"
)

func (d *PostgresDBRepo) GetItemById(ctx context.Context, id int) (models.TodoItem, error) {
	var item models.TodoItem
	result := d.DB.First(&item, id)
	if result.Error != nil {
		return item, result.Error
	}

	return item, nil
}

// GetItems returns all todo_items based on completed bool
func (d *PostgresDBRepo) GetItems(ctx context.Context, completed bool) ([]models.TodoItem, error) {
	var items []models.TodoItem

	result := d.DB.WithContext(ctx).Where("completed = ?", completed).Find(&items)

	if result.Error != nil {
		return items, result.Error
	}

	return items, nil
}

// InsertItem inserts a new todo_item
func (d *PostgresDBRepo) InsertItem(ctx context.Context, item models.TodoItem) (models.TodoItem, error) {
	result := d.DB.WithContext(ctx).Create(&item)
	if result.Error != nil {
		return item, result.Error
	}
	d.DB.Last(&item)

	return item, nil
}

// UpdateItem updates existing todo_item
func (d *PostgresDBRepo) UpdateItem(ctx context.Context, existingItem models.TodoItem, data map[string]interface{}) (models.TodoItem, error) {
	result := d.DB.WithContext(ctx).Model(&existingItem).Updates(data)
	if result.Error != nil {
		return existingItem, result.Error
	}

	return existingItem, nil
}
