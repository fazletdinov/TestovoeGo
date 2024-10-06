package app

import (
	"fmt"
	"net"

	"tasks/config"
	taskController "tasks/internal/api/grpc/controller"
	taskRepository "tasks/internal/domain/repository"
	taskService "tasks/internal/domain/service"
	taskGrpc "tasks/protogen/tasks"

	"github.com/rs/zerolog"
	"github.com/uptrace/bun"
	"google.golang.org/grpc"
)

type GRPC struct {
	log        *zerolog.Logger
	gRPCServer *grpc.Server
	Env        *config.Config
}

func NewGRPC(log *zerolog.Logger, env *config.Config, db *bun.DB) *GRPC {
	gRPCServer := grpc.NewServer()
	taskGrpc.RegisterTasksServer(
		gRPCServer,
		&taskController.TaskController{
			TaskService: taskService.NewTaskService(taskRepository.NewTaskRepository(db)),
			Env:         env,
		},
	)

	return &GRPC{
		log:        log,
		gRPCServer: gRPCServer,
		Env:        env,
	}
}

func (g *GRPC) MustRun() {
	if err := g.Run(); err != nil {
		panic(err)
	}
}

func (g *GRPC) Run() error {
	const op = "grpcapp.Run"
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", g.Env.GRPC.TasksGRPCPort))
	if err != nil {
		return fmt.Errorf("%s:%v", op, err)
	}
	g.log.Info().Msg(fmt.Sprintf("gRPC Server запущен %s", lis.Addr().String()))

	if err = g.gRPCServer.Serve(lis); err != nil {
		return fmt.Errorf("%s:%v", op, err)
	}
	return nil
}

func (g *GRPC) Stop() {
	g.log.Info().Msg(fmt.Sprintf("Остановка gRPC Server, port = %d", g.Env.GRPC.TasksGRPCPort))
	g.gRPCServer.GracefulStop()
}
