package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial value: ", i)

	// function passed by value so ival is a distinct copy of i and changes to ival do not affect i
	zeroval(i)
	fmt.Println("zeroval: ",i)


	//memory address of poiter &i
	//function deferences pointer so changes to *iptr change i 
	zeroptr(&i)
	fmt.Println("zeroptr: ", i)


	fmt.Println("pointer: ", &i)
}
