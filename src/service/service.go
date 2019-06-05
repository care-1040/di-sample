package service

import (
	. "github.com/care-1040/di-sample/src"
	"github.com/care-1040/di-sample/src/dao"
	"strconv"
)

// dao(Data Access Object)を呼び出して、dataを取得してformatする。
type TodoService struct {
}
func (s *TodoService) GetTodo(id int) string {
	dao := dao.NewTodoDao()
	return s.format(dao.Get(id))
}
func (s *TodoService) format(todo Todo) string {
	return strconv.Itoa(todo.Id) + ", " + todo.Body
}

