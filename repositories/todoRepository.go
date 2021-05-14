package repositories

import "api-echo-template/graph/model"

// NewTodoRepository :
func NewTodoRepository() *TodoRepository {
	return &TodoRepository{}
}

// TodoRepository :
type TodoRepository struct {
}

// GetTodos :
func (repo *TodoRepository) GetTodos() ([]*model.Todo, error) {
	todos := []*model.Todo{
		{ID: "72cdc181-99a4-47d0-be87-b809cb7e13f5",
			Text: "some todo", Done: true,
			User: &model.User{ID: "72cdc181-99a4-47d0-be87-b809cb7e13f5",
				Name: "Jesus"}}}
	return todos, nil
}

//SaveTodo :
func (repo *TodoRepository) SaveTodo(newTodo *model.NewTodo) (*model.Todo, error) {
	todo := model.Todo{ID: "72cdc181-99a4-47d0-be87-b809cb7e13f5",
		Text: newTodo.Text,
		Done: false,
		User: &model.User{
			ID:   "72cdc181-99a4-47d0-be87-b809cb7e13f23",
			Name: "Jesus"}}
	return &todo, nil
}
