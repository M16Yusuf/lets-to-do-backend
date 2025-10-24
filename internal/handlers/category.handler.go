package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/m16yusuf/lets-to-do/internal/models"
	"github.com/m16yusuf/lets-to-do/internal/repositories"
)

type CategoryHandler struct {
	cr *repositories.CategoryRepository
}

func NewCategoryHandler(cr *repositories.CategoryRepository) *CategoryHandler {
	return &CategoryHandler{cr: cr}
}

func (c *CategoryHandler) GetAllCategory(ctx *gin.Context) {
	categories, err := c.cr.GetAllCategory()
	if err != nil {
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
			Data: categories,
		})
	}
}

func (c *CategoryHandler) CreateCategory(ctx *gin.Context) {
	var body models.Category

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

	if err := c.cr.CreateCategory(body); err != nil {
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
			Msg:       "Category created successfully",
		})
	}
}
