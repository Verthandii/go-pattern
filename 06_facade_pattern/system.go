package main

import "fmt"

type AccountSystem struct{}

func (*AccountSystem) GetUserId(name string) (id int) {
	return 1
}

type RechargeSystem struct{}

func (*RechargeSystem) Recharge(uid, money int) {
	fmt.Printf("user id: %d, recharge money %d\n", uid, money)
}

type AuthSystem struct{}

func (*AuthSystem) Verified(uid int) bool {
	fmt.Printf("user id: %d verified\n", uid)
	return true
}
