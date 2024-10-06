package main

import (
	"os"
	"os/signal"
	"syscall"

	_ "tasks/docs"
	taskRoute "tasks/internal/api/http/route"
	"tasks/internal/app"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Gin Tasks Service
// @version         1.0
// @description     Сервис для создания Tasks.

// @contact.name   Идель Фазлетдинов
// @contact.email  fvi-it@mail.ru

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

func main() {
	app := app.App()

	env := app.Env
	go app.GRPCServer.MustRun()
	//log := app.Log

	gin := gin.Default()
	// gin.Use(logger.Logger(log))

	gin.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	taskRoute.SetupTaskRouter(gin, env, app.DB)

	gin.Run(":" + env.TasksServer.TasksPort)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	<-stop
	app.GRPCServer.Stop()
}
