package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func add(param1 int, param2 int) int{
	return param1 + param2
}

func TestAdd(t *testing.T){
	t.Run("",func(t *testing.T){
		assert.Equal(t,3, add(1,2))
	})
}

func main(){

}
