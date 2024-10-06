package route

import (
	"tasks/config"
	controller "tasks/internal/api/http/controller"

	"github.com/gin-gonic/gin"
)

func NewCreateTaskRouter(
	group *gin.RouterGroup,
	env *config.Config,
) {
	taskController := &controller.CreateTaskController{
		Env: env,
	}
	group.POST("/task", taskController.Create)
}
