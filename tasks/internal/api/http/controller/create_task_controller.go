package controller

import (
	"net/http"

	"tasks/config"
	"tasks/internal/domain/service"
	dto "tasks/internal/dto"
	"tasks/utils"

	"github.com/gin-gonic/gin"
)

type CreateTaskController struct {
	TaskService service.TaskService
	Env         *config.Config
}

// CreateTask	godoc
// @Summary		Создание Task
// @Tags        Tasks
// @Accept		json
// @Produce     json
// @Param		body	    body		dto.TaskRequest		  true	  "Создание Task"
// @Success     201  		{object}  	dto.SuccessResponse
// @Failure		400			{object}	dto.ErrorResponse
// @Failure		500			{object}	dto.ErrorResponse
// @Router      /task 		[post]
func (tc *CreateTaskController) Create(ctx *gin.Context) {
	var request dto.TaskRequest

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: "Невалидные данные"})
		return
	}

	if !utils.ValidStatus(request.Status) {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: "Статус может быть 'Выполнено' или 'Не выполнено'"})
		return
	}

	err = tc.TaskService.CreateTask(
		ctx,
		&request,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: "Internal Server error"})
		return
	}

	ctx.JSON(http.StatusCreated, dto.SuccessResponse{Message: "Task успешно создан"})
}
