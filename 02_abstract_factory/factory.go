package main

// Shop 产品族的抽象工厂
type Shop interface {
	ProduceMilkTea() MilkTea
	ProduceIceCream() IceCream
}

// YiDianDianShop 一点点 产品族
type YiDianDianShop struct{}

func NewYiDianDianShop() *YiDianDianShop {
	return &YiDianDianShop{}
}

func (*YiDianDianShop) ProduceMilkTea() MilkTea {
	return &YiDianDianMilkTea{}
}
func (*YiDianDianShop) ProduceIceCream() IceCream {
	return &YiDianDianIceCream{}
}

// CoCoShop CoCo 产品族
type CoCoShop struct{}

func NewCoCoShop() *CoCoShop {
	return &CoCoShop{}
}

func (*CoCoShop) ProduceMilkTea() MilkTea {
	return &CoCoMilkTea{}
}

func (*CoCoShop) ProduceIceCream() IceCream {
	return &CoCoIceCream{}
}
