package route

import (
	"tasks/config"
	controller "tasks/internal/api/http/controller"

	"github.com/gin-gonic/gin"
)

func NewDeleteTaskRouter(
	group *gin.RouterGroup,
	env *config.Config,
) {
	taskController := &controller.DeleteTaskController{
		Env: env,
	}

	group.DELETE("/task/:task_id", taskController.Delete)
}
