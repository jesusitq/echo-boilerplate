package repositories

import "api-echo-template/graph/model"

// ITodoRepository :
type ITodoRepository interface {
	GetTodos() ([]*model.Todo, error)
	SaveTodo(newTodo *model.NewTodo) (*model.Todo, error)
}
