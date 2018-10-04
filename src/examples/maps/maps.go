// maps aka hashes

package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["k1"]=7
	m["k2"]=13

	fmt.Println("map:", m)

	//print first value
	fmt.Println("value 1:", m["k1"])

	//length of map
	fmt.Println("length:", len(m))

	//delete key
	delete(m, "k2")
	fmt.Println("map:", m)

	//is a key present
	// optional second return value indicates if present, as we don't need the value assign it to the blank identifer _

	_, present := m["k2"]
	fmt.Println("present:", present)

	//declare and init

	n := map[string]int{"foo":1, "bar":2}
	fmt.Println("map :",n)
	
	
}
