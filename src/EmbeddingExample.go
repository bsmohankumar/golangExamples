package main

import (
	"fmt"
)
//Go doesn't support opps but we can achive same by composition under struct for inheritence
func main() {

	type Animal struct {
		animalName string
		animalOrigin string
	}

	type Bird struct {
		Animal
		birdName string
		birdColor string
		speed int
		canFly bool
	}

	bird := Bird {}
	bird.birdName = "peacock"
	bird.birdColor = "Black with color"
	bird.animalName = "Flying Bird category"
	bird.speed = 100
	bird.animalOrigin = "India"
	bird.canFly = true

	fmt.Println(bird)


}
