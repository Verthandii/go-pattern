package main

func main() {
	factories := make([]MilkTeaFactory, 0)
	factories = append(factories, NewYiDianDianFactory(), NewCoCoFactory())
	for _, factory := range factories {
		factory.ProduceMilkTea().Drink()
	}
}
