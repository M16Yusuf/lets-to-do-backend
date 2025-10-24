package repositories

import (
	"errors"
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

func (cr *CategoryRepository) UpdateCategory(id int, updates map[string]interface{}) error {
	if len(updates) == 0 {
		return errors.New("no fields to update")
	}

	result := cr.db.Model(&models.Category{}).Where("id = ?", id).Updates(updates)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
