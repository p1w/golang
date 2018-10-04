package main

import "fmt"

func main() {

	// create a zero valued array of length 5
	var a [5]int
	fmt.Println("empty:", a)

	a[4] = 100
	fmt.Println("entire array:",a)
	fmt.Println("get element: a[4] is ", a[4])

	//length of array
	fmt.Println("len:", len(a))

	//declare and initialise an array
	b := [5]int{1,2,3,4,5}
	fmt.Println("declare:",b)

	// arrays are one-dimensional but you can compose types to build multi-dimensional structures

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("twoD: ",twoD)
}

