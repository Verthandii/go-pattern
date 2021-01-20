package main

import "fmt"

// MilkTea 奶茶的抽象类
type MilkTea interface {
	Drink()
}

// YiDianDianMilkTea 一点点奶茶
type YiDianDianMilkTea struct{}

func (*YiDianDianMilkTea) Drink() {
	fmt.Println("喝一点点奶茶...")
}

// CoCoMilkTea CoCo奶茶
type CoCoMilkTea struct{}

func (*CoCoMilkTea) Drink() {
	fmt.Println("喝CoCo奶茶...")
}
