package controller

import (
	"net/http"
	"strconv"

	"tasks/config"
	"tasks/internal/domain/service"
	dto "tasks/internal/dto"

	"github.com/gin-gonic/gin"
)

type GetTaskController struct {
	taskService service.TaskService
	Env         *config.Config
}

// GetTask	   godoc
// @Summary	   Получение Task
// @Tags       Tasks
// @Accept	   json
// @Produce    json
// @Param	   task_id		      path		    string		          	true		"Task ID"
// @Success    200  		      {object}  	dto.TaskResponse
// @Failure	   500			      {object}	    dto.ErrorResponse
// @Router     /task/{task_id}    [get]
func (tc *GetTaskController) Fetch(ctx *gin.Context) {
	taskID := ctx.Param("task_id")
	taskID64, err := strconv.ParseInt(taskID, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: "Невалидные данные"})
		return
	}

	task, err := tc.taskService.GetTask(
		ctx,
		taskID64,
	)
	if err != nil {
		ctx.JSON(http.StatusNotFound, dto.ErrorResponse{Message: "Task не найден"})
		return
	}

	ctx.JSON(http.StatusOK, task)
}

// ListTasks godoc
//
// @Summary		Получение списока Task
// @Tags	    Tasks
// @Accept	    json
// @Produce		json
// @Param	    limit			query				int		true	"limit"
// @Param	    offset			query				int		true	"offset"
// @Success		200	{array}		dto.TaskResponse
// @Failure		400	{object}	dto.ErrorResponse
// @Failure		404	{object}	dto.ErrorResponse
// @Failure		500	{object}	dto.ErrorResponse
// @Router	    /tasks 			[get]
func (tc *GetTaskController) Fetchs(ctx *gin.Context) {
	limit, err := strconv.Atoi(ctx.Query("limit"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: "Невалидные данные"})
		return
	}
	offset, err := strconv.Atoi(ctx.Query("offset"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: "Невалидные данные"})
		return
	}

	tasks, err := tc.taskService.GetTasks(ctx, limit, offset)
	if err != nil {
		ctx.JSON(http.StatusNotFound, dto.ErrorResponse{Message: "Tasks не найден"})
		return
	}

	ctx.JSON(http.StatusOK, tasks)
}
