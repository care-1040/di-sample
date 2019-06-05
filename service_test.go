package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTodoService(t *testing.T) {
	service := TodoService{}
	result := service.GetTodo(1)
	expected := "1, aaa"
	assert.Equal(t, result, expected)

}

