package controller

import (
	"context"
	"strconv"

	"tasks/config"
	taskService "tasks/internal/domain/service"
	taskGrpc "tasks/protogen/tasks"
	"tasks/utils"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TaskController struct {
	taskGrpc.UnimplementedTasksServer
	TaskService taskService.TaskService
	Env         *config.Config
}

func (tc *TaskController) GetTasks(
	ctx context.Context,
	taskRequest *taskGrpc.GetTasksRequest,
) (*taskGrpc.GetTasksResponse, error) {

	tasks, err := tc.TaskService.GetTasks(
		ctx,
		int(taskRequest.GetLimit()),
		int(taskRequest.GetOffset()),
	)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal server error")
	}

	tasksResponse := make([]*taskGrpc.Task, 0, taskRequest.Limit)
	for _, task := range *tasks {
		tasksResponse = append(tasksResponse, &taskGrpc.Task{
			TaskId:      strconv.Itoa(int(task.ID)),
			Title:       task.Title,
			Description: task.Description,
			Status:      task.Status,
			CreatedAt:   task.CreatedAt.GoString(),
			UpdatedAt:   task.UpdatedAt.GoString(),
		})
	}
	return &taskGrpc.GetTasksResponse{
		Tasks: tasksResponse,
	}, nil
}

func (tc *TaskController) UpdateTaskStatus(
	ctx context.Context,
	taskRequest *taskGrpc.UpdateTaskRequest,
) (*taskGrpc.UpdateTaskResponse, error) {

	if taskRequest.GetTaskId() == "" {
		return nil, status.Error(codes.InvalidArgument, "поле task_id обязательно")
	}

	if taskRequest.GetStatus() == "" {
		return nil, status.Error(codes.InvalidArgument, "поле status обязательно")
	}

	taskID64, err := strconv.ParseInt(taskRequest.GetTaskId(), 0, 64)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "ошибка при парсинге taskID")
	}

	if !utils.ValidStatus(taskRequest.GetStatus()) {
		return nil, status.Error(codes.InvalidArgument, "Статус может быть 'Выполнено' или 'Не выполнено'")
	}

	_, err = tc.TaskService.GetTask(
		ctx,
		taskID64,
	)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}

	err = tc.TaskService.UpdateTask(
		ctx,
		taskID64,
		taskRequest.GetStatus(),
	)

	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &taskGrpc.UpdateTaskResponse{
		Message: "Обновление прошло успешно",
	}, nil
}
