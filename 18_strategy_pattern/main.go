package main

import "fmt"

func main() {
	sumMachine := NewMachine(&SumStrategy{})
	fmt.Println(sumMachine.Calc(10, 5))
	minusMachine := NewMachine(&MinusStrategy{})
	fmt.Println(minusMachine.Calc(10, 5))
	// Output:
	// 15
	// 5
}
