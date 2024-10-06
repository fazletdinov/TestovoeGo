package service

import (
	"context"
	"tasks/internal/domain/repository"
	"tasks/internal/dto"
	"tasks/internal/models"
)

type taskService struct {
	taskRepository repository.TaskRepository
}

func NewTaskService(taskRepository repository.TaskRepository) TaskService {
	return &taskService{
		taskRepository: taskRepository,
	}
}

func (ts *taskService) CreateTask(
	ctx context.Context,
	body *dto.TaskRequest,
) error {
	task := models.Task{
		Titile:      body.Title,
		Description: body.Description,
		Status:      body.Status,
	}
	return ts.taskRepository.CreateTask(ctx, &task)
}

func (ts *taskService) GetTask(
	ctx context.Context,
	taskID int,
) (*dto.TaskResponse, error) {
	task, err := ts.taskRepository.GetTask(ctx, taskID)
	if err != nil {
		return nil, err
	}
	taskResponse := dto.TaskResponse{
		ID:          task.ID,
		Title:       task.Titile,
		Description: task.Description,
		Status:      task.Status,
		CreatedAt:   task.CreatedAt,
		UpdatedAt:   task.UpdatedAt,
	}
	return &taskResponse, nil
}

func (ts *taskService) GetTasks(
	ctx context.Context,
	limit int,
	offset int,
) (*[]dto.TaskResponse, error) {
	tasks, err := ts.taskRepository.GetTasks(ctx, limit, offset)
	if err != nil {
		return nil, err
	}
	tasksResponse := make([]dto.TaskResponse, limit)
	for _, task := range *tasks {
		tasksResponse = append(tasksResponse, dto.TaskResponse{
			ID:          task.ID,
			Title:       task.Titile,
			Description: task.Description,
			Status:      task.Status,
			CreatedAt:   task.CreatedAt,
			UpdatedAt:   task.UpdatedAt,
		})
	}
	return &tasksResponse, nil
}

func (ts *taskService) UpdateTask(
	ctx context.Context,
	taskID int64,
	status string,
) error {
	return ts.taskRepository.UpdateTask(ctx, taskID, status)
}

func (ts *taskService) DeleteTask(
	ctx context.Context,
	taskID int64,
) error {
	return ts.taskRepository.DeleteTask(ctx, taskID)
}
