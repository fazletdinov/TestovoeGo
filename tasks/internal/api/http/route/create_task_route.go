package route

import (
	"tasks/config"
	controller "tasks/internal/api/http/controller"
	"tasks/internal/domain/repository"
	"tasks/internal/domain/service"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

func NewCreateTaskRouter(
	group *gin.RouterGroup,
	env *config.Config,
	db *bun.DB,
) {
	taskRepository := repository.NewTaskRepository(db)
	taskService := service.NewTaskService(taskRepository)
	taskController := &controller.CreateTaskController{
		TaskService: taskService,
		Env:         env,
	}
	group.POST("/task", taskController.Create)
}
