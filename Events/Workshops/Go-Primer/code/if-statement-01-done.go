package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// START OMIT
func max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}

// END OMIT

func Test_max(t *testing.T) {
	assert.Equal(t, 1, max(0, 1))
}

func main() {
	var tests []testing.InternalTest
	tests = append(tests, testing.InternalTest{Name: "Test_max", F: Test_max})
	testing.Main(func(pat, str string) (bool, error) { return true, nil }, tests, nil, nil)
}
