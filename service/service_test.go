package service

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 現状serviceのテストにmysqlが必要
func TestTodoService1(t *testing.T) {
	service := TodoService{}
	result := service.GetTodo(1)
	expected := "1, aaa"
	assert.Equal(t, result, expected)
}

func TestTodoService2(t *testing.T) {
	service := TodoService{}
	result := service.GetTodo(2)
	expected := "neet"
	assert.Equal(t, result, expected)
}
