package main

import "fmt"

type AbstractShape interface {
	Drawing()
}

type ColorImplementer interface {
	Color() string
}

type RedColorImplementer struct{}

func (*RedColorImplementer) Color() string {
	return "红色"
}

func RedColor() *RedColorImplementer {
	return &RedColorImplementer{}
}

type BlueColorImplementer struct{}

func BlueColor() *BlueColorImplementer {
	return &BlueColorImplementer{}
}

func (*BlueColorImplementer) Color() string {
	return "蓝色"
}

type Rectangle struct {
	color ColorImplementer
}

func NewRectangle(color ColorImplementer) *Rectangle {
	return &Rectangle{color: color}
}

func (r *Rectangle) Drawing() {
	fmt.Println(r.color.Color() + "的长方形")
}

type Circle struct {
	color ColorImplementer
}

func NewCircle(color ColorImplementer) *Circle {
	return &Circle{color: color}
}

func (c *Circle) Drawing() {
	fmt.Println(c.color.Color() + "的圆形")
}
