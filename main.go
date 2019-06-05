package main

import "fmt"

func main() {
	service := TodoService{}
	fmt.Println(service.GetTodo(1))
}