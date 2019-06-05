package main

import (
	"fmt"
	. "github.com/care-1040/di-sample/src/service"
	. "github.com/care-1040/di-sample/src/dao"
)

func main() {
	dao := NewTodoDao()
	service := TodoService{Dao: dao}
	fmt.Println(service.GetTodo(1))
}