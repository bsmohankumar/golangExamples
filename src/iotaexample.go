package main
import (
	"fmt"
)

const (
	_ = iota
	KB = 1<<(10*iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

const (
	isAdmin = 1 << iota
	isSales
	isFinansials
	isIt
	isHomeGaurd

)

func main() {
 fileSize := 4000000000.
 fmt.Printf("%.2fGB", fileSize/GB)

 var roles byte = isAdmin | isSales | isFinansials | isHomeGaurd
    fmt.Printf("%b\n", isAdmin)
	fmt.Printf("%b\n", isSales)
	fmt.Printf("%b\n", isFinansials)
	fmt.Printf("%b\n", isHomeGaurd)
	fmt.Printf("%b\n", roles)
	fmt.Printf("is Admin ? %v\n", isAdmin & roles == isAdmin)


}
