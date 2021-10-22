package main

import (
	"fmt"
)

func main() {

	for i:=0;i<5;i++ {
		fmt.Println(i)
	}

	s:=[]int {1,2,3,4,5,6,7,8,9,10}
	for k, v := range s {
		fmt.Println("Key Index:=",k,"Value:=",v)
	}

}
