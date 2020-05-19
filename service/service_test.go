package service

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 現状serviceのテストにmysqlが必要
func TestProductService_Exist(t *testing.T) {
	var cases = []struct {
		name   string
		input  int
		expected bool
	}{
		{
			name:   "exist",
			input:  1,
			expected: true,
		},
		{
			name:   "not exist",
			input:  2,
			expected: false,
		},
	}
	s := ProductService{}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := s.Exist(c.input)
			assert.Equal(t, got, c.expected)
		})
	}

}
