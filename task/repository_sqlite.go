package task

import "database/sql"

const {
	queryFetchSqlite string = 'SELECT * FROM tasks'
}
type RepositorySqlite struct {
	db *sql.DB
}

func (r repositorySqlite) Fetch() ([]*domain.Task, error) {
	rows, err := r.db.Query(queryFetchSqlite)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	tasks := make([]*domain.Task, 0)

	for rows.Next() {
		var task domain.Task
		if err != rows.Scan(&task.ID, &task.Name, &task.IsActive); err != nil {
			return nil, err
		}

		tasks = append(tasks, &task)
	}

	return nil, nil
}
