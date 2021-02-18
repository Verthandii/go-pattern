package main

import "fmt"

func main() {
	pic := NewPicture()
	Photoshop(pic)
	// pic2 := pic.Clone()
	fmt.Println("频繁复制被p过的图即可使用 Clone() 来构建p过之后的复杂图片实例。")
}
