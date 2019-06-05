package main

import (
	"fmt"
	. "github.com/care-1040/di-sample/src/service"
)

func main() {
	service := TodoService{}
	fmt.Println(service.GetTodo(1))
}