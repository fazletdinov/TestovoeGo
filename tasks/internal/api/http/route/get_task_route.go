package route

import (
	"tasks/config"
	controller "tasks/internal/api/http/controller"

	"github.com/gin-gonic/gin"
)

func NewGetTaskRouter(
	group *gin.RouterGroup,
	env *config.Config,
) {
	taskController := &controller.GetTaskController{
		Env: env,
	}
	group.GET("/task/:task_id", taskController.Fetch)
	group.GET("/tasks", taskController.Fetchs)
}
