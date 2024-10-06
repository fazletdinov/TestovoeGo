package controller

import (
	"net/http"
	"strconv"

	"tasks/config"
	"tasks/internal/domain/service"
	dto "tasks/internal/dto"

	"github.com/gin-gonic/gin"
)

type UpdateTaskController struct {
	taskService service.TaskService
	Env         *config.Config
}

// UpdateTask   godoc
// @Summary     Обновление Task
// @Tags        Tasks
// @Accept		json
// @Produce     json
// @Param	    task_id			path		string					    true    "Task ID"
// @Param		body		    body		dto.UpdateTaskRequest	    true	"Для обновления Task"
// @Success     200  		    {object}  	dto.SuccessResponse
// @Failure	  	400			    {object}	dto.ErrorResponse
// @Failure	  	401			    {object}	dto.ErrorResponse
// @Failure	  	500			    {object}	dto.ErrorResponse
// @Router      /task/{task_id} [put]
func (utc *UpdateTaskController) Update(ctx *gin.Context) {
	taskID := ctx.Param("task_id")
	taskID64, err := strconv.ParseInt(taskID, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: "Невалидные данные"})
		return
	}

	_, err = utc.taskService.GetTask(
		ctx,
		taskID64,
	)
	if err != nil {
		ctx.JSON(http.StatusNotFound, dto.ErrorResponse{Message: "Task не найден"})
		return
	}

	var taskRequest dto.UpdateTaskRequest

	if err = ctx.ShouldBindJSON(&taskRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: "Невалидные  данные"})
		return
	}
	err = utc.taskService.UpdateTask(
		ctx,
		taskID64,
		taskRequest.Status,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: "Internal Server error"})
		return
	}

	ctx.JSON(http.StatusOK, dto.SuccessResponse{Message: "Обновление прошло успешно"})

}
