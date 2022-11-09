// package name is always at top like in Java.
// Go is structured in modules and packages.
package main

// formatted input / output - package
// unused imports will not compile unlike other programming languages
import "fmt"

// main function like C, programme starts here.
func main() {
	// difference, no delimiter at the end of the line with semicolon
	// delimiter seems optional
	// format style for Go global for all linters unlike JavaScript
	// method naming different
	fmt.Println("Hello Universe!")
	// linebreak with \n , %d is for numbers base 10
	fmt.Printf("%d\n", 2022)
	fmt.Println("End line!")
}

// Compilation of Go Source happens in memory and will be executed from there.