package _8_phoneCall

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_solution(t *testing.T) {
	assert.Equal(t, 14, solution(3, 1, 2, 20))

	assert.Equal(t, 1, solution(2, 2, 1, 2))
	assert.Equal(t, 11, solution(10, 1, 2, 22))
	assert.Equal(t, 14, solution(2, 2, 1, 24))
	assert.Equal(t, 3, solution(1, 2, 1, 6))

	assert.Equal(t, 0, solution(10, 10, 10, 8))
}
