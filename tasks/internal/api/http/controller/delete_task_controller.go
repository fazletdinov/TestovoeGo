package controller

import (
	"net/http"
	"strconv"

	"tasks/config"
	"tasks/internal/domain/service"
	dto "tasks/internal/dto"

	"github.com/gin-gonic/gin"
)

type DeleteTaskController struct {
	taskService service.TaskService
	Env         *config.Config
}

// DeleteTask	godoc
// @Summary		Удаление Task
// @Tags        Tasks
// @Accept		json
// @Produce     json
// @Param	    task_id			path		string		          		true 	"Task ID"
// @Success     204  		    {object}  	dto.SuccessResponse
// @Failure		401			    {object}	dto.ErrorResponse
// @Failure		500			    {object}	dto.ErrorResponse
// @Router      /task/{task_id} [delete]
func (dtc *DeleteTaskController) Delete(ctx *gin.Context) {
	taskID := ctx.Param("task_id")
	taskID64, err := strconv.ParseInt(taskID, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: "Невалидные данные"})
		return
	}

	_, err = dtc.taskService.GetTask(
		ctx,
		taskID64,
	)
	if err != nil {
		ctx.JSON(http.StatusNotFound, dto.ErrorResponse{Message: "Task не найден"})
		return
	}

	if err = dtc.taskService.DeleteTask(
		ctx,
		taskID64,
	); err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: "Internal Server error"})
		return
	}

	ctx.JSON(http.StatusNoContent, dto.SuccessResponse{Message: "Task успешно удален"})

}
