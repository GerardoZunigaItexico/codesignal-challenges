package main

import "fmt"

func main(){
	fmt.Println(centuryFromYear(1905))
	fmt.Println(centuryFromYear(1700))
}

func centuryFromYear(year int) int {
	m := year%100
	c := (year-m)/100
	if m > 0{
		c = c+1
	}
	return c
}