package domain

type (
	Task struct {
		ID string 'json:"id"'
		Name string 'json:"name"'
		IsActive bool 'json:"is_active'
	}

	TaskRepository interface {
		Fetch() ([]*Task, error)
	}

	TaskService interface {
		Fetch() ([]*Task, error)
	}
)