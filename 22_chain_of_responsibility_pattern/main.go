package main

func main() {
	handler1 := NewRequestHandler1()
	handler2 := NewRequestHandler2()

	handler1.SetNext(handler2)
	//handler2.SetNext(handler1) 形成死循环
	handler1.HandleRequest("request")

	// output:
	// handler1...request
	// handler2...request
}
