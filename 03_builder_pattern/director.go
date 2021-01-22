package main

// KFCWaiter KFC 服务员 在此设计模式中充当 Director 的身份
type KFCWaiter struct {
	setMenuBuilder KFCSetMenuBuilder
}

func NewKFCWaiter() *KFCWaiter {
	return &KFCWaiter{}
}

func (k *KFCWaiter) SetKFCSetMenu(setMenuBuilder KFCSetMenuBuilder) {
	k.setMenuBuilder = setMenuBuilder
}

func (k *KFCWaiter) Construct() KFCSetMenu {
	// build ...
	k.setMenuBuilder.BuilderDrink()
	k.setMenuBuilder.BuilderEat()
	k.setMenuBuilder.BuilderTableware()

	return k.setMenuBuilder.GetKFCSetMenu()
}
