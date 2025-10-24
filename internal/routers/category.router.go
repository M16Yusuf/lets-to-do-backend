package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/m16yusuf/lets-to-do/internal/handlers"
	"github.com/m16yusuf/lets-to-do/internal/repositories"
	"gorm.io/gorm"
)

func InitCategoryRouter(router *gin.Engine, db *gorm.DB) {
	categoryRouter := router.Group("/api/categories")
	categoryRepository := repositories.NewCategoryRepository(db)
	categoryHandler := handlers.NewCategoryHandler(categoryRepository)

	categoryRouter.GET("", categoryHandler.GetAllCategory)
	categoryRouter.POST("", categoryHandler.CreateCategory)
	categoryRouter.PUT("/:id", categoryHandler.UpdateCategory)
	categoryRouter.DELETE("/:id", categoryHandler.DeleteCategory)
}
