package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/m16yusuf/lets-to-do/internal/middleware"
	"github.com/m16yusuf/lets-to-do/internal/models"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()
	router.Use(middleware.CORSMiddleware)

	// setup routing
	InitTodoRouter(router, db)
	InitCategoryRouter(router, db)

	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, models.Response{
			IsSuccess: false,
			Code:      http.StatusNotFound,
			Msg:       "Page not found!",
		})
	})

	return router
}
