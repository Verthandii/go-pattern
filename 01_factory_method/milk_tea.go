package main

import "fmt"

// MilkTea 奶茶的抽象类
type MilkTea interface {
	Drink()
}

// DiamondMilkTea 珍珠奶茶
type DiamondMilkTea struct {
}

func (d *DiamondMilkTea) Drink() {
	fmt.Println("喝珍珠奶茶...")
}

// PuddingMilkTea 布丁奶茶
type PuddingMilkTea struct {
}

func (p *PuddingMilkTea) Drink() {
	fmt.Println("喝布丁奶茶...")
}
