package service

import (
	"github.com/ha-t2/di-sample/dao"
	. "github.com/ha-t2/di-sample/model"
	"strconv"
)

// dao(Data Access Object)を呼び出して、todoを取得してformatする。
type TodoService struct {
}

func (s *TodoService) GetTodo(id int) string {
	dao := dao.NewTodoDao()
	todo := dao.Get(id)
	if todo.Body == "" {
		return "neet"
	} else {
		return s.format(todo)
	}
}
func (s *TodoService) format(todo Todo) string {
	return strconv.Itoa(todo.Id) + ", " + todo.Body
}
