package service

import (
	. "github.com/care-1040/di-sample/src"
	"github.com/care-1040/di-sample/src/dao"
	"strconv"
)

// dao(Data Access Object)を呼び出して、todoを取得してformatする。
type TodoService struct {
	Dao dao.TodoDaoInterface
}
func (s *TodoService) GetTodo(id int) string {
	todo := s.Dao.Get(id)
	if todo.Body == "" {
		return "neet"
	} else {
		return s.format(todo)
	}
}
func (s *TodoService) format(todo Todo) string {
	return strconv.Itoa(todo.Id) + ", " + todo.Body
}

