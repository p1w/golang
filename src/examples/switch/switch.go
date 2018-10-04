package main

import "fmt"
import "time"

func main() {
	i:= 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}


	// use commas to have multiple expressions in the same case 
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("weekend")
	default:
		fmt.Println("weekday")
	}

	// switch without an expression is alternative to if/else
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("am")
	default:
		fmt.Println("pm")
	}

	// a type switch compares types instead of values
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("bool")
		case int:
			fmt.Println("int")
		default:
			fmt.Printf("don't know type %T\n",t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
