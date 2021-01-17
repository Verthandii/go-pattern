package main

// MilkTeaFactory 奶茶的抽象工厂
type MilkTeaFactory interface {
	// Produce 这个方法就是 工厂方法
	Produce() MilkTea
}

// YiDianDianFactory 一点点奶茶店
type YiDianDianFactory struct{}

func NewYiDianDianFactory() *YiDianDianFactory {
	return &YiDianDianFactory{}
}

func (y *YiDianDianFactory) Produce() MilkTea {
	return &DiamondMilkTea{}
}

// CoCoFactory CoCo奶茶店
type CoCoFactory struct{}

func NewCoCoFactory() *CoCoFactory {
	return &CoCoFactory{}
}

func (c *CoCoFactory) Produce() MilkTea {
	return &PuddingMilkTea{}
}
