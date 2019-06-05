package main

import (
	"fmt"
	"github.com/ha-t2/di-sample/service"
)

func main() {
	service := service.TodoService{}
	fmt.Println(service.GetTodo(1))
}
