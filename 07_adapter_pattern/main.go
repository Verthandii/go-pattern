package main

import "fmt"

func main() {
	adaptee := NewAdaptee()
	target := NewAdapter(adaptee)
	fmt.Println(target.Request())
}
