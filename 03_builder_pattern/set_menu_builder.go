package main

import "fmt"

// KFCSetMenuBuilder KFC套餐生成器
type KFCSetMenuBuilder interface {
	BuilderDrink()
	BuilderEat()
	BuilderTableware()

	GetKFCSetMenu() KFCSetMenu
}

// KFCCoupleSetMenuBuilder KFC 情侣套餐生成器
type KFCCoupleSetMenuBuilder struct{}

func (*KFCCoupleSetMenuBuilder) BuilderDrink() {
	fmt.Println("准备两个大可...")
}

func (*KFCCoupleSetMenuBuilder) BuilderEat() {
	fmt.Println("准备六对鸡翅...")
}

func (*KFCCoupleSetMenuBuilder) BuilderTableware() {
	fmt.Println("准备两双手套和两个吸管...")
}

func (*KFCCoupleSetMenuBuilder) GetKFCSetMenu() KFCSetMenu {
	return &KFCCoupleSetMenu{}
}

// KFCFamilySetMenuBuilder KFC 全家桶套餐生成器
type KFCFamilySetMenuBuilder struct{}

func (*KFCFamilySetMenuBuilder) BuilderDrink() {
	fmt.Println("准备三个大可...")
}

func (*KFCFamilySetMenuBuilder) BuilderEat() {
	fmt.Println("准备九对鸡翅...")
}

func (*KFCFamilySetMenuBuilder) BuilderTableware() {
	fmt.Println("准备三双手套和三个吸管...")
}

func (*KFCFamilySetMenuBuilder) GetKFCSetMenu() KFCSetMenu {
	return &KFCFamilySetMenu{}
}
