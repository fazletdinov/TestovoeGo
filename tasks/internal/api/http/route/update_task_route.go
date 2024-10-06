package route

import (
	"tasks/config"
	controller "tasks/internal/api/http/controller"

	"github.com/gin-gonic/gin"
)

func NewUpdateTaskRouter(
	group *gin.RouterGroup,
	env *config.Config) {
	taskController := &controller.UpdateTaskController{
		Env: env,
	}
	group.PUT("/task/:task_id", taskController.Update)
}
