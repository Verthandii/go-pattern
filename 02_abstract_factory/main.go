package main

func main() {
	factories := make([]Shop, 0)
	factories = append(factories, NewYiDianDianShop(), NewCoCoShop())
	for _, factory := range factories {
		factory.ProduceMilkTea().Drink()
		factory.ProduceIceCream().Eat()
	}
}
