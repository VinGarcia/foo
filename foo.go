package foo

import (
	"fmt"

	"github.com/vingarcia/bar"
)

func Main() {
	fmt.Printf("In foo package rigit now\n")

	fmt.Println("printing a Bar struct:", bar.Bar{})
}
