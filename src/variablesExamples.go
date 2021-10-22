package main

import (
	"fmt"
	"strconv"
)

const a int16 = 127
const b1 int16 = 16
func main() {

	var i int
	i = 27
	var j int = 42
	//k:=99
	var m float64 = 7.5
	var s1 string = strconv.Itoa(i)
	fmt.Printf("%v \t %T \n",i, i)
	fmt.Printf("%v \t %T \n",j, j)
	//fmt.Printf("%v \t %T \n",k, k)
	k:= m
	fmt.Printf("%v \t %T \n",k, k)
	fmt.Printf("%v \t %T \n",s1, s1)

	n:= 1 == 1
	m1:= 1 == 2

	fmt.Printf("%v \t %T \n",n, n)
	fmt.Printf("%v \t %T \n",m1, m1)

	t:=10
	r:=3
	fmt.Println(t&r)
	fmt.Println(t|r)
	fmt.Println(t^r)
	fmt.Println(t&^r)

	var str string = "this is first string"
	str1:= "this is second string"

	fmt.Printf("%v, %T",str,str)
	fmt.Printf("%v, %T",str1,str1)

	fmt.Printf("\n %v, %T",string(str[2]),str)
	fmt.Printf("\n %v, %T",str1+str,str)
	b:=[]byte(str)
	fmt.Printf("\n %v, %T",b,b)

	for i := 1; i < len(b); i++ {
		fmt.Printf("\n %v, %T",string(b[i]),b)
	}

	const a int = 16

	fmt.Printf("\n %v, %T",a,a)
	fmt.Printf("\n %v, %T",b1,b1)

}
