package repositories

import (
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
