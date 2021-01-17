package main

func main() {
	milkTeas := make([]MilkTea, 0)

	yidiandian := NewYiDianDianFactory()
	coco := NewCoCoFactory()

	milkTeas = append(milkTeas, yidiandian.Produce(), coco.Produce())
	for _, milkTea := range milkTeas {
		milkTea.Drink()
	}
}
