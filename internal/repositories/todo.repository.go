package repositories

import (
	"context"
	"time"

	"github.com/m16yusuf/lets-to-do/internal/models"
	"gorm.io/gorm"
)

type TodoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) *TodoRepository {
	return &TodoRepository{db: db}
}

func (td *TodoRepository) CreateTodo(cntxt context.Context, body models.Todo) error {
	newTodo := models.Todo{
		Title:       body.Title,
		Description: body.Description,
		Category_id: body.Category_id,
		Priority:    body.Priority,
		DueDate:     body.DueDate,
		CreatedAt:   time.Now(),
		UpdatedAt:   nil,
	}
	if err := td.db.WithContext(cntxt).Create(&newTodo).Error; err != nil {
		return err
	}
	return nil
}
