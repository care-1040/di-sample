package main

import "strconv"

type TodoService struct {
}

func (s *TodoService) GetTodo(id int) string {
	dao := NewTodoDao()
	return s.format(dao.Get(id))
}
func (s *TodoService) format(todo Todo) string {
	return strconv.Itoa(todo.Id) + ", " + todo.Body
}

