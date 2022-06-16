package domain

type CommandStatus struct {
	Success      bool   `json:"success"`
	ErrorMessage string `json:"errorMessage"`
}

func Success() CommandStatus {
	return CommandStatus{Success: true}
}

func Failure(errorMessage string) CommandStatus {
	return CommandStatus{Success: false, ErrorMessage: errorMessage}
}

type SelectTodosQuery struct{}

type SelectTodosQueryResult struct {
	Todos []Todo `json:"todos"`
}

type SelectTodosQueryHandler func(q SelectTodosQuery) SelectTodosQueryResult
