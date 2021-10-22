package main

import
(
	"fmt"
)

func main() {

	if true {
		fmt.Println("Inside True block")
	}

	if 2!=2 {
		fmt.Println("Inside True block")
	} else{
		fmt.Println("Inside False block")
	}

	switch 2 {
	case 1: fmt.Println("Inside 1 block")
	case 2: fmt.Println("Inside 2 block")
	default:
		fmt.Println("Inside default block")

	}
}
