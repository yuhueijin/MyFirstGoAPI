package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	type args struct {
		model *Model
	}

	testCases := []struct {
		name string
		args args
		err  error
	}{
		{
			name: "should add items",
			args: args{
				model: &Model{
					Name: "test item",
				},
			},
			err: nil,
		},
	}

	targetService := NewProductService()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := targetService.Add(tc.args.model.Name)

			assert.Equal(t, tc.err, err)
		})
	}
}

func TestRemove(t *testing.T) {
	type args struct {
		id int
	}

	testCases := []struct {
		name string
		args args
		err  error
	}{
		{
			name: "should remove items",
			args: args{
				id: 1,
			},
			err: nil,
		},
	}

	s := NewProductService()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := s.Remove(tc.args.id)

			assert.Equal(t, tc.err, err)
		})
	}
}

func TestGetAll(t *testing.T) {
	testCases := []struct {
		name     string
		expected []Model
		err      error
	}{
		{
			name:     "should get all items",
			expected: []Model{},
			err:      nil,
		},
	}

	s := NewProductService()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := s.GetAll()

			assert.Equal(t, tc.expected, actual)
			assert.Equal(t, tc.err, err)
		})
	}
}
