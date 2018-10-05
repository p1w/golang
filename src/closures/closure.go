package main

import "fmt"

//this functions returns another function, defined anonymously in the body. the returned function closes over variable i to form a closure

func intSeq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main() {

	// this function value captures its own i value
	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())


	//new function - new value of i
	newInts := intSeq()
	fmt.Println(newInts())
}
