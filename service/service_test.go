package service

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 現状serviceのテストにmysqlが必要
func TestMoneyService1(t *testing.T) {
	s := ProductService{}
	result := s.GetPriceWithTax(1)
	expected := 110
	assert.Equal(t, result, expected)
}

