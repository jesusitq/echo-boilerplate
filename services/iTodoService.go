package services

import "api-echo-template/graph/model"

// ITodoService :
type ITodoService interface {
	Get() ([]*model.Todo, error)
	Save(newTodo *model.NewTodo) (*model.Todo, error)
}
