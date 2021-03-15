package main

func main() {
	camera := NewFilterCamera(NewBeautyCamera(NewCamera()))
	camera.Shoot()
}
