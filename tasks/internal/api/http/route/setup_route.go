package route

import (
	"tasks/config"

	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
)

func SetupTaskRouter(
	gin *gin.Engine,
	env *config.Config,
	db *bun.DB,
) {
	publicRouter := gin.Group("/api/v1")
	NewGetTaskRouter(publicRouter, env, db)

	protectedRouter := gin.Group("/api/v1")
	// protectedRouter.Use(middleware.JwtAuthMiddleware(env.JWTConfig.PathPublicKey))
	NewCreateTaskRouter(protectedRouter, env, db)
	NewDeleteTaskRouter(protectedRouter, env, db)
	NewUpdateTaskRouter(protectedRouter, env, db)
}
