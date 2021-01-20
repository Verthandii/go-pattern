package main

// MilkTeaFactory 奶茶的抽象工厂
type MilkTeaFactory interface {
	Produce() MilkTea
}

// DiamondMilkTeaFactory 珍珠奶茶工厂
type DiamondMilkTeaFactory struct{}

func NewDiamondMilkTeaFactory() *DiamondMilkTeaFactory {
	return &DiamondMilkTeaFactory{}
}

func (*DiamondMilkTeaFactory) Produce() MilkTea {
	return &DiamondMilkTea{}
}

// PuddingMilkTeaFactory 布丁奶茶工厂
type PuddingMilkTeaFactory struct{}

func NewPuddingMilkTeaFactory() *PuddingMilkTeaFactory {
	return &PuddingMilkTeaFactory{}
}

func (*PuddingMilkTeaFactory) Produce() MilkTea {
	return &PuddingMilkTea{}
}
