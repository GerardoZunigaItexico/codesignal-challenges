package _3_is_infinite_process

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_solution01(t *testing.T) {
	assert.Equal(t, false, solution01(2, 6))
	assert.Equal(t, true, solution01(2, 3))
	assert.Equal(t, true, solution01(0, 1))
	assert.Equal(t, true, solution01(3, 1))
}
func Test_solution02(t *testing.T) {
	assert.Equal(t, false, solution02(2, 6))
	assert.Equal(t, true, solution02(2, 3))
	assert.Equal(t, true, solution02(0, 1))
	assert.Equal(t, true, solution02(3, 1))
}

func Test_solution03(t *testing.T) {
	assert.Equal(t, false, solution03(2, 6))
	assert.Equal(t, true, solution03(2, 3))
	assert.Equal(t, true, solution03(0, 1))
	assert.Equal(t, true, solution03(3, 1))
}
