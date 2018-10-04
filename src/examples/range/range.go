//range iterates over a data type

package main

import "fmt"

func main() {
	//create a slice
	nums := []int{2,3,4}
	sum := 0

	//the first value is the index - here we don't use it so assign to _
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum: ", sum)

	//this example uses the index
	for index, num := range nums {
		if num == 3 {
			fmt.Println("index:", index)
		}
	}

	fruit := map[string]string{"a": "apple", "b": "banana"}
	for k,v := range fruit {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for key := range fruit {
		fmt.Println("key:",key)
	}

	//range on strings iterates over unicode code points. first value is the index, second is the char
	for i, c := range "golang" {
		fmt.Println(i, c)
	}

	

}
