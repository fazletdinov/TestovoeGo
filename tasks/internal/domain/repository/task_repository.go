package repository

import (
	"context"
	"time"

	"tasks/internal/models"

	"github.com/uptrace/bun"
)

type taskRepository struct {
	database *bun.DB
}

func NewTaskRepository(db *bun.DB) TaskRepository {
	return &taskRepository{
		database: db,
	}
}

func (tr *taskRepository) CreateTask(
	ctx context.Context,
	body *models.Task,
) error {
	err := tr.database.NewInsert().Model(body).Scan(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (tr *taskRepository) GetTask(
	ctx context.Context,
	taskID int64,
) (*models.Task, error) {
	task := new(models.Task)
	err := tr.database.NewSelect().Model(task).Where("id = ?", taskID).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (tr *taskRepository) GetTasks(
	ctx context.Context,
	limit int,
	offset int,
) (*[]models.Task, error) {
	var tasks []models.Task
	err := tr.database.NewSelect().Model(&tasks).Offset(offset).Limit(limit).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return &tasks, nil
}

func (tr *taskRepository) UpdateTask(
	ctx context.Context,
	taskID int64,
	status string,
) error {
	task := &models.Task{ID: taskID, Status: status, UpdatedAt: time.Now()}
	_, err := tr.database.NewUpdate().Model(task).OmitZero().WherePK().Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (tr *taskRepository) DeleteTask(
	ctx context.Context,
	taskID int64,
) error {
	task := &models.Task{ID: taskID}
	_, err := tr.database.NewDelete().Model(task).WherePK().Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}
