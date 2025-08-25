package mysql

import (
	"context"
	"database/sql"
	
	"patricego/usecases/domain"
)

type TaskRepo struct {
	DB *sql.DB
}

func NewTaskRepository(db *sql.DB) domain.TaskRepository {
	return &TaskRepo{DB: db}
}

func (r *TaskRepo) CreateTask(ctx context.Context, task domain.Task) (int64, error) {
	query := `INSERT INTO tasks(resource_id, name, description, created_time, updated_time, status) VALUES(?, ?, ?, ?, ?, ?)`
	result, err := r.DB.ExecContext(ctx, query, task.ResourceID, task.Name, task.Description, task.CreatedTime, task.UpdatedTime, task.Status)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func (r *TaskRepo) GetAllTasks(ctx context.Context) ([]domain.Task, error) {
	query := `SELECT id, resource_id, name, description, created_time, updated_time, status FROM tasks WHERE status = 'A'`
	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []domain.Task
	for rows.Next() {
		var task domain.Task
		err := rows.Scan(&task.Id, &task.ResourceID, &task.Name, &task.Description, &task.CreatedTime, &task.UpdatedTime, &task.Status)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (r *TaskRepo) GetTaskByResourceID(ctx context.Context, resourceID string) (domain.Task, error) {
	query := `SELECT id, resource_id, name, description, created_time, updated_time, status FROM tasks WHERE resource_id = ? AND status = 'A'`
	var task domain.Task
	err := r.DB.QueryRowContext(ctx, query, resourceID).Scan(&task.Id, &task.ResourceID, &task.Name, &task.Description, &task.CreatedTime, &task.UpdatedTime, &task.Status)
	return task, err
}

func (r *TaskRepo) UpdateTask(ctx context.Context, task domain.Task) error {
	query := `UPDATE tasks SET name = ?, description = ?, status = ?, updated_time = ? WHERE resource_id = ?`
_, err := r.DB.ExecContext(ctx, query, task.Name, task.Description, task.Status, task.UpdatedTime, task.ResourceID)
	return err
}

func (r *TaskRepo) UpdateTaskStatus(ctx context.Context, resourceID string, status string) error {
	query := `UPDATE tasks SET status = ?, updated_time = NOW() WHERE resource_id = ?`
	_, err := r.DB.ExecContext(ctx, query, status, resourceID)
	return err
}


func (r *TaskRepo) DeleteTask(ctx context.Context, resourceID string) error {
	query := `UPDATE tasks SET status = 'D' WHERE resource_id = ? AND status = 'A'`
	_, err := r.DB.ExecContext(ctx, query, resourceID)
	return err
}
