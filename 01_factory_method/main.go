package main

func main() {
	factories := make([]MilkTeaFactory, 0)
	factories = append(factories, NewDiamondMilkTeaFactory(), NewPuddingMilkTeaFactory())
	for _, factory := range factories {
		factory.Produce().Drink()
	}
}
