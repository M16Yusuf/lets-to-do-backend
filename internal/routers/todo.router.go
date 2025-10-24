package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/m16yusuf/lets-to-do/internal/handlers"
	"github.com/m16yusuf/lets-to-do/internal/repositories"
	"gorm.io/gorm"
)

func InitTodoRouter(router *gin.Engine, db *gorm.DB) {
	todoRouter := router.Group("/api/todos")
	todoRepository := repositories.NewTodoRepository(db)
	todohandler := handlers.NewTodoHandler(todoRepository)

	todoRouter.GET("", todohandler.GetAllAndFilter)
	todoRouter.POST("", todohandler.CreateTodo)
	todoRouter.GET("/:id", todohandler.GetDetailTodo)
	todoRouter.PUT("/:id", todohandler.UpdateTodo)
	todoRouter.DELETE("/:id", todohandler.DeleteTodo)
	todoRouter.PATCH("/:id/complete", todohandler.ToggleComplete)
}
