package main
import (
	"fmt"
)
func main() {

	type Doctor  struct {

		doctorId int
		doctorName string
		doctorAge int
		spouse []string
	}

	aDoctor := Doctor{
		doctorId : 1,
		doctorName: "Guru Raghavendra",
		doctorAge: 1,
		spouse: []string{"Sarswathi"},
	}

	fmt.Println(aDoctor)
	fmt.Println("Doctor Name : ",aDoctor.doctorName)
	fmt.Println("Doctor ID : ",aDoctor.doctorId)
	fmt.Println("Doctor Age : ",aDoctor.doctorAge)
	fmt.Println("Doctor Spouse : ",aDoctor.spouse)

}
