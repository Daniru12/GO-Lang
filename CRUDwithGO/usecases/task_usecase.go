package usecases

import (
	"context"
	"errors"
	"strings"
	"patricego/usecases/domain"
)

type TaskUsecase struct {
	TaskRepo domain.TaskRepository
}

func NewTaskUsecase(repo domain.TaskRepository) domain.TaskUsecase {
	return &TaskUsecase{TaskRepo: repo}
}


func (uc *TaskUsecase) CreateTask(ctx context.Context, task domain.Task) (int64, error) {

	if strings.TrimSpace(task.Name) == "" {
		return 0, errors.New("task name cannot be empty")
	}
	if len(task.Name) > 100 {
		return 0, errors.New("task name is too long (max 100 chars)")
	}
	if strings.TrimSpace(task.Description) == "" {
		return 0, errors.New("task description cannot be empty")
	}


	if task.Status == "" {
		task.Status = "A" 
	}

	return uc.TaskRepo.CreateTask(ctx, task)
}


func (uc *TaskUsecase) GetAllTasks(ctx context.Context) ([]domain.Task, error) {
	return uc.TaskRepo.GetAllTasks(ctx)
}

func (uc *TaskUsecase) GetTaskByResourceID(ctx context.Context, resourceID string) (domain.Task, error) {
	if strings.TrimSpace(resourceID) == "" {
		return domain.Task{}, errors.New("resourceID cannot be empty")
	}
	return uc.TaskRepo.GetTaskByResourceID(ctx, resourceID)
}


func (uc *TaskUsecase) UpdateTask(ctx context.Context, task domain.Task) error {
	if task.Id == 0 {
		return errors.New("task ID is required for update")
	}
	if strings.TrimSpace(task.Name) == "" {
		return errors.New("task name cannot be empty")
	}
	return uc.TaskRepo.UpdateTask(ctx, task)
}


func (uc *TaskUsecase) UpdateTaskStatus(ctx context.Context, resourceID string) error {
	if strings.TrimSpace(resourceID) == "" {
		return errors.New("resourceID cannot be empty")
	}
	return uc.TaskRepo.UpdateTaskStatus(ctx, resourceID, "D")
}
