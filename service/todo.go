package service

import "github.com/thanawatpetchuen/ginstructure/model"

type TodoService struct{}

func NewTodoService() *TodoService {
	return &TodoService{}
}

func (s *TodoService) GetTodos() model.TodoResponse {
	response := model.TodoResponse{
		Todos: []string{"test", "test2"},
	}
	return response
}
