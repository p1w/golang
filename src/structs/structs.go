//structs are typed collections of fields. used to group data into records

package main

import "fmt"

// person struct has name and age fields
type person struct {
	name string
	age int
}

func main() {

	//create a new struct
	fmt.Println(person{"Bob",20})

	//can name fields when initialising
	fmt.Println(person{name: "Alice", age:32})

	//omitted fields are zero valued
	fmt.Println(person{name: "Franko"})

	//an & prefix yeilds a pointer to the struct
	fmt.Println(&person{name: "Dave", age: 55})

	//access struct fields with a dot
	s := person{"philip",38}
	fmt.Println(s.name)

	//can use dots with struct pointers.
	sp := &s
	fmt.Println(sp.age)

	//structs are mutable
	sp.age = 99
	fmt.Println(sp.age)
}
