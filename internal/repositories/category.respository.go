package repositories

import (
	"time"

	"github.com/m16yusuf/lets-to-do/internal/models"
	"gorm.io/gorm"
)

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (cr *CategoryRepository) GetAllCategory() ([]models.Category, error) {
	var categories []models.Category

	if err := cr.db.Find(&categories).Error; err != nil {
		return nil, err
	}

	return categories, nil
}

func (cr *CategoryRepository) CreateCategory(body models.Category) error {
	newCategory := models.Category{
		Name:      body.Name,
		Color:     body.Color,
		CreatedAt: time.Now(),
	}

	if err := cr.db.Create(&newCategory).Error; err != nil {
		return err
	}

	return nil
}
