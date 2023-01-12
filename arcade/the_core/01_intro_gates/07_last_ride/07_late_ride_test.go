package _7_last_ride

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_solution(t *testing.T) {
	assert.Equal(t, 4, solution(240))
	assert.Equal(t, 14, solution(808))
	assert.Equal(t, 19, solution(1439))
	assert.Equal(t, 0, solution(0))
	assert.Equal(t, 8, solution(8))

}
