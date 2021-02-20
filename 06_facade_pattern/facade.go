package main

type FacadeSystem interface {
	UserRecharge() int
}

type GateWay struct{}

func NewGateWay() *GateWay {
	return &GateWay{}
}

func (*GateWay) UserRecharge(name string, money int) int {
	accountSys, rechargeSys, authSys := &AccountSystem{}, &RechargeSystem{}, &AuthSystem{}
	// 从用户系统得到用户id
	uid := accountSys.GetUserId(name)
	// 从认证系统验证此用户是否合法
	if !authSys.Verified(uid) {
		return -1
	}
	// 从充值系统中充值
	rechargeSys.Recharge(uid, money)
	return 0
}
