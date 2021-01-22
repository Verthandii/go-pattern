package main

import "fmt"

func main() {
	waiter := NewKFCWaiter()
	waiter.SetKFCSetMenu(&KFCCoupleSetMenuBuilder{})
	coupleSetMenu := waiter.Construct()
	coupleSetMenu.EnjoyIt()

	fmt.Println()

	waiter.SetKFCSetMenu(&KFCFamilySetMenuBuilder{})
	familySetMenu := waiter.Construct()
	familySetMenu.EnjoyIt()
}
