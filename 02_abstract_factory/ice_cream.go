package main

import "fmt"

type IceCream interface {
	Eat()
}

type YiDianDianIceCream struct{}

func (*YiDianDianIceCream) Eat() {
	fmt.Println("吃一点点冰淇淋")
}

type CoCoIceCream struct{}

func (*CoCoIceCream) Eat() {
	fmt.Println("吃CoCo冰淇淋")
}
