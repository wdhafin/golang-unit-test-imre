package learn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd_WithTestTable(t *testing.T) {
	testCases := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{
			name:     "Negative and Negative",
			a:        -1,
			b:        -1,
			expected: -2,
		},
		{
			name:     "Negative and Positive",
			a:        -1,
			b:        1,
			expected: 0,
		},
		{
			name:     "Positive and Positive",
			a:        1,
			b:        1,
			expected: 2,
		},
		{
			name:     "Positive and Negative",
			a:        1,
			b:        -1,
			expected: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			c := Add(tc.a, tc.b)
			assert.Equal(t, tc.expected, c)
		})
	}
}

func TestAbsolute(t *testing.T) {
	t.Run("Negative Test Case", func(t *testing.T) {
		c := Absolute(-5)
		assert.Equal(t, 5, c)
	})
	t.Run("Positive Test Case", func(t *testing.T) {
		c := Absolute(5)
		assert.Equal(t, 5, c)
	})
}

// func TestAdd(t *testing.T) {
// 	t.Run("Negative add negative", func(t *testing.T) {
// 		c := Add(-1, -2)
// 		assert.Equal(t, -3, c)
// 	})
// 	t.Run("Positive add negative", func(t *testing.T) {
// 		c := Add(1, -2)
// 		assert.Equal(t, -1, c)
// 	})
// }

func TestAdd(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Positive and Positive",
			args: args{
				a: 1,
				b: 2,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}
