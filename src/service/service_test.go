package service

import (
	"github.com/care-1040/di-sample/src/mock_dao"
	"github.com/stretchr/testify/assert"
	"github.com/care-1040/di-sample/src"
	"github.com/golang/mock/gomock"
	"testing"
)

// 現状serviceのテストにmysqlが必要
func TestTodoService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	dao := mock_main.NewMockTodoDaoInterface(ctrl)
	dao.EXPECT().Get(1).Return(src.Todo{Id: 1, Body: "aaa"})

	service := TodoService{dao}
	result := service.GetTodo(1)
	expected := "1, aaa"
	assert.Equal(t, result, expected)
}

func TestTodoService2(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	dao := mock_main.NewMockTodoDaoInterface(ctrl)
	dao.EXPECT().Get(2).Return(src.Todo{Id: 2, Body: ""})

	service := TodoService{dao}
	result := service.GetTodo(2)
	expected := "neet"
	assert.Equal(t, result, expected)
}

