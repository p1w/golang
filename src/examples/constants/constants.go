package main

import "fmt"
import "math"


//A const statement can appear anywhere a var statement can.
//Constant expressions perform arithmetic with arbitrary precision.

const s string = "constant"

func main() {
	fmt.Println(s)

	const n = 50000000

	const d = 3e20 / n
	fmt.Println(d)
// a constant has no type until it's given one such as by an explicit cast
	fmt.Println(int64(d))
//or by using in a context which requires one e.g. Sin expects a float64
	fmt.Println(math.Sin(n))
}
