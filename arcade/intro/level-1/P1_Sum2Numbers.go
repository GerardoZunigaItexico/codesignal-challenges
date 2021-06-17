package main

import "fmt"

func main(){
	fmt.Println(add(5,7))
}

func add(param1 int, param2 int) int{
	return param1 + param2
}
