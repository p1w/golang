package main

import "fmt"

func main(){

	var a string = "intial"
	fmt.Println(a)

	var b, c int = 1,2
	fmt.Println(b,c)

	//Go will infer the type of initialized variables.
	var d = true
	fmt.Println(d)

	// vars declared without initialization are zero-valued
	var e int
	fmt.Println(e)

	// syntax := is shorthand for declaring and initilaising
	f := "short"
	fmt.Println(f)
}
