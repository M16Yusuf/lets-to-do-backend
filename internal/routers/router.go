package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/m16yusuf/lets-to-do/internal/models"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	// setup routing
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, models.Response{
			IsSuccess: true,
			Code:      http.StatusOK,
			Msg:       "pong",
		})
	})

	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, models.Response{
			IsSuccess: false,
			Code:      http.StatusNotFound,
			Msg:       "Page not found!",
		})
	})

	return router
}
