package repository

import (
	"context"

	"tasks/internal/models"
)

type TaskRepository interface {
	CreateTask(ctx context.Context, body *models.Task) error
	GetTask(ctx context.Context, taskID int64) (*models.Task, error)
	GetTasks(ctx context.Context, limit int, offset int) (*[]models.Task, error)
	UpdateTask(ctx context.Context, taskID int64, status string) error
	DeleteTask(ctx context.Context, taskID int64) error
}
