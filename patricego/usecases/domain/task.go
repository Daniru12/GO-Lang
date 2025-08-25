package domain

import (
	"context"
	"time"
)

type Task struct {
	Id          int64     `json:"id"`
	ResourceID  string    `json:"resource_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedTime time.Time    `json:"created_time"`
	UpdatedTime time.Time    `json:"updated_time"`
	Status      string    `json:"status"`
}

type TaskRepository interface {
	CreateTask(ctx context.Context, task Task) (int64, error)
	GetAllTasks(ctx context.Context) ([]Task, error)
	GetTaskByResourceID(ctx context.Context, resourceID string) (Task, error)
	UpdateTask(ctx context.Context, task Task) error
	
	UpdateTaskStatus(ctx context.Context, resourceID string, status string) error
}

type TaskUsecase interface {
	CreateTask(ctx context.Context, task Task) (int64, error)
	GetAllTasks(ctx context.Context) ([]Task, error)
	GetTaskByResourceID(ctx context.Context, resourceID string) (Task, error)
	UpdateTask(ctx context.Context, task Task) error
	
	UpdateTaskStatus(ctx context.Context, resourceID string) error
}
