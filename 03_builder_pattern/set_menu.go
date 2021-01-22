package main

import "fmt"

type KFCSetMenu interface {
	EnjoyIt()
}

type KFCCoupleSetMenu struct{}

func (*KFCCoupleSetMenu) EnjoyIt() {
	fmt.Println("和对象一起吃 KFC 情侣套餐真是太好啦~")
}

type KFCFamilySetMenu struct{}

func (*KFCFamilySetMenu) EnjoyIt() {
	fmt.Println("和爸妈一起吃 KFC 全家桶套餐真是太好啦~")
}
