package main

// Import package math with subpackage rand
// combined imports
import (
	"fmt"
	"math/rand"
)


func main(){
	// prints a random number, half open interval between 0 and 100, so it will be between 0 and 99
	fmt.Println("Hello new number: ", rand.Intn(100))
}