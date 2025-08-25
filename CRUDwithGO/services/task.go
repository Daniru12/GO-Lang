package services

import (
	"context"
	"patricego/usecases/domain"
)

type TaskService struct {
	TaskUsecase domain.TaskUsecase
}

func NewTaskService(usecase domain.TaskUsecase) *TaskService {
	return &TaskService{
		TaskUsecase: usecase,
	}
}

func (s *TaskService) CreateTask(ctx context.Context, task domain.Task) (int64, error) {
	return s.TaskUsecase.CreateTask(ctx, task)
}

func (s *TaskService) GetAllTasks(ctx context.Context) ([]domain.Task, error) {
	return s.TaskUsecase.GetAllTasks(ctx)
}

func (s *TaskService) GetTaskByResourceID(ctx context.Context, resourceID string) (domain.Task, error) {
	return s.TaskUsecase.GetTaskByResourceID(ctx, resourceID)
}

func (s *TaskService) UpdateTask(ctx context.Context, task domain.Task) error {
	return s.TaskUsecase.UpdateTask(ctx, task)
}

func (s *TaskService) UpdateTaskStatus(ctx context.Context, resourceID string) error {
	return s.TaskUsecase.UpdateTaskStatus(ctx, resourceID)
}




