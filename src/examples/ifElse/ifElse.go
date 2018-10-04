package main

import "fmt"

func main() {

	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}
	// a statement can precede conditionals; any vars declared in the statement are available in all branches.

	if num := 9; num < 0 {
		fmt.Println (num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	// num is undefined here
	fmt.Println(num, "outside")
}

