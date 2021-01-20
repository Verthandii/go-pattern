package main

// MilkTeaFactory 奶茶的抽象工厂
type MilkTeaFactory interface {
	ProduceMilkTea() MilkTea
}

// YiDianDianFactory 一点点奶茶店
type YiDianDianFactory struct{}

func NewYiDianDianFactory() *YiDianDianFactory {
	return &YiDianDianFactory{}
}

func (*YiDianDianFactory) ProduceMilkTea() MilkTea {
	return &YiDianDianMilkTea{}
}

// CoCoFactory CoCo奶茶店
type CoCoFactory struct{}

func NewCoCoFactory() *CoCoFactory {
	return &CoCoFactory{}
}

func (*CoCoFactory) ProduceMilkTea() MilkTea {
	return &CoCoMilkTea{}
}
