package unittest

import (
	"gopractise/server"
	"testing"

	assert "github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {

	expectedAnswer := 3

	actualAnswer := server.Add(1, 2)

	assert.Equal(t, expectedAnswer, actualAnswer)
}
