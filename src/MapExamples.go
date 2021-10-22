package main

import(
	"fmt"
)

func main() {
	statePopulations := make(map[string]int) //using make method to define a map
	statePopulations = map[string]int{
		"Karnataka":700000,
		"Hydrabad":800000,
		"Hariyana":89,
		"Panjab":90,
	}
	fmt.Println(statePopulations)
	statePopulations["Delhi"]=45
	fmt.Println(statePopulations)
	duplicateMap := statePopulations
	delete(statePopulations,"Panjab")
	fmt.Println("length statePopulations ", len(statePopulations))
	fmt.Println("length duplicateMap ", len(duplicateMap))
	fmt.Println(statePopulations)
	pop , ok := statePopulations["Panjab"]
	fmt.Println(pop, ok)
	pop , ok1 := statePopulations["Delhi"]
	fmt.Println(pop, ok1)
	fmt.Println("length ", len(statePopulations))

	m:= map[[3]int]string{}
	fmt.Println(m)
}
