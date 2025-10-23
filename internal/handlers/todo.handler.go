package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/m16yusuf/lets-to-do/internal/models"
	"github.com/m16yusuf/lets-to-do/internal/repositories"
)

type Todohandler struct {
	tr *repositories.TodoRepository
}

func NewTodoHandler(tr *repositories.TodoRepository) *Todohandler {
	return &Todohandler{tr: tr}
}

func (t *Todohandler) CreateTodo(ctx *gin.Context) {
	var body models.Todo

	if err := ctx.ShouldBindJSON(&body); err != nil {
		log.Println("failed binding data \nCause :", err)
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Response: models.Response{
				IsSuccess: false,
				Code:      http.StatusInternalServerError,
			},
			Err: "Internal server error",
		})
		return
	}

	if err := t.tr.CreateTodo(ctx.Request.Context(), body); err != nil {
		log.Println("failed execute repositories \n Cause : ", err)
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Response: models.Response{
				IsSuccess: false,
				Code:      http.StatusInternalServerError,
			},
			Err: "Internal server error",
		})
	} else {
		ctx.JSON(http.StatusOK, models.Response{
			IsSuccess: true,
			Code:      http.StatusOK,
			Msg:       "Todo created successfully",
		})
	}
}
