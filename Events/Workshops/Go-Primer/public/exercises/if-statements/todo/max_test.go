package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_max(t *testing.T) {
	assert.Equal(t, 1, max(0, 1))
}
