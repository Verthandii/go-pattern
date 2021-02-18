package main

import "fmt"

type Cloneable interface {
	Clone() Cloneable
}

type Picture struct {
	Name   string
	Data   []byte
	Length int
	Width  int
}

func NewPicture() *Picture {
	return &Picture{}
}

func (p *Picture) Clone() *Picture {
	c := *p
	return &c
}

func Photoshop(pic *Picture) {
	fmt.Println("p图完成！")
}
