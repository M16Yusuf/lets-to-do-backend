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

	todoRouter.POST("", todohandler.CreateTodo)
}
