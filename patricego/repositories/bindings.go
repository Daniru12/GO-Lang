package repositories

import(
	"database/sql"
	"patricego/repositories/mysql"
	"patricego/usecases/domain"
)

type Repositories struct{
	TaskRepo domain.TaskRepository
}

func NewRepositories(db *sql.DB) *Repositories{
	return &Repositories{
		TaskRepo: mysql.NewTaskRepository(db),
	}
}
