package main

import "fmt"

func main() {

	i:=1
	for i <= 3 {
		fmt.Println(i)
		i++
	}
	for j := 7; j<=9; j++ {
		fmt.Println(j)
	}

	// loop without condition will repeat until break
	for {
		fmt.Println("loop")
		break
	}

	//continue to next iteration of the loop
	for n := 0; n<=5; n++ {
		if n%2 != 0 {
			continue

		}
		fmt.Println(n)
	}


}
