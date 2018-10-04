package main

import "fmt"

func main() {

	s := make([]string,3)
	fmt.Println("empty slice:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("slice setup:", s)
	fmt.Println("get 2nd:", s[1])

	fmt.Println("len:", len(s))

	//append to a slice
	s = append(s, "d")
	s = append(s, "e","f")
	fmt.Println("appended", s)

	//make a copy of a slice - slices can be different size either will stop once full or will have empty values at end
	c := make([]string,len(s))
	copy (c, s)
	fmt.Println("copy:", c)


	//get a part of a slice
	l := s[2:5]
	fmt.Println("part of slice:", l)

	//slice upto(but excluding)
	fmt.Println("upto s[:5]", s[:5])

	//slice from (and including)

	fmt.Println("upto s[2:]", s[2:])

	//declare and init in one line
	t := []string{"g","h","i"}
	fmt.Println("declare and init", t)

}
