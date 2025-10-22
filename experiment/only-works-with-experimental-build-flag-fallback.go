//go:build !experimental_feature

package experiment

import "fmt"

func IsActive() {
	fmt.Println("IS INACTIVE!")
}
