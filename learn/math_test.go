package learn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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

func TestAdd(t *testing.T) {
	t.Run("Negative add negative", func(t *testing.T) {
		c := Add(-1, -2)
		assert.Equal(t, -3, c)
	})
	t.Run("Positive add negative", func(t *testing.T) {
		c := Add(1, -2)
		assert.Equal(t, -1, c)
	})
}
