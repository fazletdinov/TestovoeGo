package service

import (
	"context"

	"tasks/internal/dto"
)

type TaskService interface {
	CreateTask(ctx context.Context, body *dto.TaskRequest) error
	GetTask(ctx context.Context, taskID int64) (*dto.TaskResponse, error)
	GetTasks(ctx context.Context, limit int, offset int) (*[]dto.TaskResponse, error)
	UpdateTask(ctx context.Context, taskID int64, status string) error
	DeleteTask(ctx context.Context, taskID int64) error
}
