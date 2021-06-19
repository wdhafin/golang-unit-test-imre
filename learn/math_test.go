package learn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAbsolute(t *testing.T) {
	c := Absolute(-5)
	assert.Equal(t, 5, c)
}

func TestAbsolute_WithPositive(t *testing.T) {
	c := Absolute(5)
	assert.Equal(t, 5, c)
}
