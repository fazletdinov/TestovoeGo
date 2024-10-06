package route

import (
	"tasks/config"

	"github.com/gin-gonic/gin"
)

func SetupTaskRouter(
	gin *gin.Engine,
	env *config.Config,
) {
	publicRouter := gin.Group("/api/v1")
	NewGetTaskRouter(publicRouter, env)

	protectedRouter := gin.Group("/api/v1")
	// protectedRouter.Use(middleware.JwtAuthMiddleware(env.JWTConfig.PathPublicKey))
	NewCreateTaskRouter(protectedRouter, env)
	NewDeleteTaskRouter(protectedRouter, env)
	NewUpdateTaskRouter(protectedRouter, env)
}
