package main

// Target 是使用者想用的接口
type Target interface {
	Request() string
}

// Adaptee 是被适配的接口，也就是我们自己原有的接口
type Adaptee interface {
	SpecificRequest() string
}

func NewAdaptee() Adaptee {
	return &AdapteeImpl{}
}

type AdapteeImpl struct{}

func (*AdapteeImpl) SpecificRequest() string {
	return "adaptee method"
}

func NewAdapter(adaptee Adaptee) Target {
	return &Adapter{
		Adaptee: adaptee,
	}
}

// Adapter 适配器，将我们原有的接口 Adaptee 转换为使用者需要的接口 Target
type Adapter struct {
	Adaptee
}

func (a *Adapter) Request() string {
	return a.SpecificRequest()
}
