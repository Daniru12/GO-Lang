package usecases

import (
	"context"
	"patricego/usecases/domain"
)

type TaskUsecase struct {
	TaskRepo domain.TaskRepository
}

func NewTaskUsecase(repo domain.TaskRepository) domain.TaskUsecase {
	return &TaskUsecase{TaskRepo: repo}
}

func (uc *TaskUsecase) CreateTask(ctx context.Context, task domain.Task) (int64, error) {
	return uc.TaskRepo.CreateTask(ctx, task)
}

func (uc *TaskUsecase) GetAllTasks(ctx context.Context) ([]domain.Task, error) {
	return uc.TaskRepo.GetAllTasks(ctx)
}

func (uc *TaskUsecase) GetTaskByResourceID(ctx context.Context, resourceID string) (domain.Task, error) {
	return uc.TaskRepo.GetTaskByResourceID(ctx, resourceID)
}

func (uc *TaskUsecase) UpdateTask(ctx context.Context, task domain.Task) error {
	return uc.TaskRepo.UpdateTask(ctx, task)
}

func (uc *TaskUsecase) UpdateTaskStatus(ctx context.Context, resourceID string) error {
	return uc.TaskRepo.UpdateTaskStatus(ctx, resourceID, "D")
}




