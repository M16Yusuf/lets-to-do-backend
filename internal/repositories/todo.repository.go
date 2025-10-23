package repositories

import (
	"fmt"
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

func (td *TodoRepository) CreateTodo(body models.Todo) error {
	newTodo := models.Todo{
		Title:       body.Title,
		Description: body.Description,
		Category_id: body.Category_id,
		Priority:    body.Priority,
		DueDate:     body.DueDate,
		CreatedAt:   time.Now(),
		UpdatedAt:   nil,
	}
	if err := td.db.Create(&newTodo).Error; err != nil {
		return err
	}
	return nil
}

func (td *TodoRepository) GetAllTodos(search string, page, limit int, sort string) ([]models.TodoList, int64, error) {
	var todos []models.TodoList
	var total int64

	query := td.db.Model(&models.Todo{})

	// Filter pencarian jika ada query param ?q=
	if search != "" {
		searchPattern := fmt.Sprintf("%%%s%%", search)
		query = query.Where("title ILIKE ? OR description ILIKE ?", searchPattern, searchPattern)
	}

	// Hitung total data sebelum pagination
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Urutan sort (default DESC = terbaru)
	order := "created_at DESC"
	if sort == "asc" {
		order = "created_at ASC"
	}

	// Pagination
	offset := (page - 1) * limit
	if offset < 0 {
		offset = 0
	}

	err := query.Order(order).
		Limit(limit).
		Offset(offset).
		Find(&todos).Error

	if err != nil {
		return nil, 0, err
	}

	return todos, total, nil
}
