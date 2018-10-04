//return multiple avlues - used to return the result and error values
package main

import "fmt"

func vals() (int, int) {
	return 3, 7
}

func main() {

	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	// use _, if you don't want all the values
	_, c :=vals()
	fmt.Println(c)
}
