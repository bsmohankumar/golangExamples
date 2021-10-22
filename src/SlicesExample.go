package main

import (
	"fmt"
)

func main() {

	a:= []int {1,2,3,4,5,6,7,8,9,10}
	fmt.Println(a)
	b := a[:]
	fmt.Println(b)
	c := a[:3]
	fmt.Println(c)
	d := a[3:]
	fmt.Println(d)
	e := a[3:6]
	fmt.Println(e)

}
