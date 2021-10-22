package main
import(
	"fmt"
	"reflect"
)
func main() {

	type  Animal struct{
		Name string `required max :"4"`
		Orgin string
	}

	animal := Animal {
		Name: "Cow",
		Orgin: "Desi",

	}

	t:=reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)
	fmt.Println(animal)
}
