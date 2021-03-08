package main

func main() {
	subject := NewRealSubject()
	proxy := NewProxy(subject)
	proxy.Do()
}
