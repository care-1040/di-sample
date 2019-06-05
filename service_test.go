package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 現状serviceのテストにmysqlが必要
func TestTodoService(t *testing.T) {
	service := TodoService{}
	result := service.GetTodo(1)
	expected := "1, aaa"
	assert.Equal(t, result, expected)
}

