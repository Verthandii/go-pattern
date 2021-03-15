package main

import "fmt"

func main() {
	factory := GetFlyweightMapFactory()
	user1Map1 := factory.GetFlyweightMap(1)
	user1Map2 := factory.GetFlyweightMap(1)
	user2Map1 := factory.GetFlyweightMap(2)
	fmt.Println(user1Map1 == user1Map2) // true
	fmt.Println(user1Map1 == user2Map1) // false
}
