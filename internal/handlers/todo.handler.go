package handlers

import (
	"log"
	"net/http"
	"strconv"

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

	if err := t.tr.CreateTodo(body); err != nil {
		log.Println("failed execute repositories \n Cause : ", err)
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Response: models.Response{
				IsSuccess: false,
				Code:      http.StatusInternalServerError,
			},
			Err: "Internal server error",
		})
		return
	} else {
		ctx.JSON(http.StatusOK, models.Response{
			IsSuccess: true,
			Code:      http.StatusOK,
			Msg:       "Todo created successfully",
		})
	}
}

func (t *Todohandler) GetAllAndFilter(ctx *gin.Context) {
	search := ctx.Query("query")
	limit := 10
	pageStr := ctx.DefaultQuery("page", "1")
	page, _ := strconv.Atoi(pageStr)
	sort := ctx.DefaultQuery("sort", "desc")

	todos, total, err := t.tr.GetAllTodos(search, page, limit, sort)
	if err != nil {
		log.Println("failed execute repositories \n Cause : ", err)
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Response: models.Response{
				IsSuccess: false,
				Code:      http.StatusInternalServerError,
			},
			Err: "Internal server error",
		})
		return
	} else {
		ctx.JSON(http.StatusOK, models.ResponseData{
			Response: models.Response{
				IsSuccess: true,
				Code:      http.StatusOK,
			},
			Data:  todos,
			Page:  page,
			Limit: limit,
			Total: int(total),
		})
	}
}

func (t *Todohandler) GetDetailTodo(ctx *gin.Context) {
	todoID, _ := strconv.Atoi(ctx.Param("id"))

	todo, err := t.tr.GetDetailTodo(todoID)
	if err != nil {
		log.Println("failed execute repositories \n Cause : ", err)
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Response: models.Response{
				IsSuccess: false,
				Code:      http.StatusInternalServerError,
			},
			Err: "Internal server error",
		})
		return
	} else {
		ctx.JSON(http.StatusOK, models.ResponseData{
			Response: models.Response{
				IsSuccess: true,
				Code:      http.StatusOK,
			},
			Data: todo,
		})
	}
}
