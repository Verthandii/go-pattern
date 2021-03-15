package main

func main() {
	NewCircle(RedColor()).Drawing()
	NewCircle(BlueColor()).Drawing()
	NewRectangle(RedColor()).Drawing()
	NewRectangle(BlueColor()).Drawing()

	// output:
	// 红色的圆形
	// 蓝色的圆形
	// 红色的长方形
	// 蓝色的长方形
}
