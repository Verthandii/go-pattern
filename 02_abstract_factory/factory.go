package main

// MilkTeaFactory 奶茶的抽象工厂
type MilkTeaFactory interface {
	ProduceMilkTea() MilkTea
	ProduceIceCream() IceCream
}

// YiDianDianFactory 一点点奶茶店
type YiDianDianFactory struct{}

func NewYiDianDianFactory() *YiDianDianFactory {
	return &YiDianDianFactory{}
}

func (*YiDianDianFactory) ProduceMilkTea() MilkTea {
	return &YiDianDianMilkTea{}
}
func (*YiDianDianFactory) ProduceIceCream() IceCream {
	return &YiDianDianIceCream{}
}

// CoCoFactory CoCo奶茶店
type CoCoFactory struct{}

func NewCoCoFactory() *CoCoFactory {
	return &CoCoFactory{}
}

func (*CoCoFactory) ProduceMilkTea() MilkTea {
	return &CoCoMilkTea{}
}

func (*CoCoFactory) ProduceIceCream() IceCream {
	return &CoCoIceCream{}
}
