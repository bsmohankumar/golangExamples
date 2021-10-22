package main

import (
	"fmt"
)

func main() {

	grades :=[3] int {97,98,99}
	grades1 :=[...] int {97,98,99,100}
	fmt.Printf("Grades %v\n",grades)
	fmt.Printf("Grades1 %v\n",grades1)
	var students [3] string
	fmt.Printf("students %v\n",students)
	students[0] = "Bhuvan BM"
	students[2] = "Beena CB"
	students[1] = "Mohan Kumar BS"
	fmt.Printf("students %v\n",students)
	fmt.Printf("Number of students %v\n",len(students))

	var matrix [3][3]int
	matrix[0]=[3]int{1,2,3}
	matrix[1]=[3]int{4,5,6}
	matrix[2]=[3]int{7,8,9}
	fmt.Printf("Matrix %v\n",matrix)



}
