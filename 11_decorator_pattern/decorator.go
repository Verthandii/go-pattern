package main

import "fmt"

type ICamera interface {
	Shoot()
}

type Camera struct{}

func NewCamera() *Camera {
	return &Camera{}
}

func (*Camera) Shoot() {
	fmt.Println("拍了一张照片")
}

type BeautyCamera struct {
	camera ICamera
}

func NewBeautyCamera(camera ICamera) *BeautyCamera {
	return &BeautyCamera{camera}
}

func (bc *BeautyCamera) Shoot() {
	fmt.Println("开启美颜")
	bc.camera.Shoot()
}

type FilterCamera struct {
	camera ICamera
}

func NewFilterCamera(camera ICamera) *FilterCamera {
	return &FilterCamera{camera}
}

func (bc *FilterCamera) Shoot() {
	fmt.Println("开启滤镜")
	bc.camera.Shoot()
}
