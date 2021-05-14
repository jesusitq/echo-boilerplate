package services

import (
	"api-echo-template/graph/model"
	"api-echo-template/repositories"
)

// NewTodoService :
func NewTodoService(repository *repositories.TodoRepository) *TodoService {
	return &TodoService{TodoRepository: repository}
}

// TodoService :
type TodoService struct {
	TodoRepository repositories.ITodoRepository
}

// Get :
func (service *TodoService) Get() ([]*model.Todo, error) {

	// insert here your bussines logic ...

	return service.TodoRepository.GetTodos()
}

// Save :
func (service *TodoService) Save(newTodo *model.NewTodo) (*model.Todo, error) {
	// insert here your bussines logic ...
	return service.TodoRepository.SaveTodo(newTodo)
}
