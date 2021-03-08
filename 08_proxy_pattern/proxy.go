package main

import "fmt"

type Subject interface {
	Do()
}

type RealSubject struct{}

func NewRealSubject() *RealSubject {
	return &RealSubject{}
}

func (*RealSubject) Do() {
	fmt.Println("object do...")
}

type Proxy struct {
	subject Subject
}

func NewProxy(subject Subject) Subject {
	return &Proxy{
		subject: subject,
	}
}

func (*Proxy) before() {
	fmt.Println("before...")
}

func (*Proxy) after() {
	fmt.Println("after...")
}

func (p *Proxy) Do() {
	p.before()
	p.subject.Do()
	p.after()
}
