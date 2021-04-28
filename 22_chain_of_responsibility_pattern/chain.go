package main

import "fmt"

type Request string

type Handler interface {
	SetNext(Handler)
	Next() Handler
	HandleRequest(Request)
}

type Handler1 struct {
	next Handler
}

func NewRequestHandler1() *Handler1 {
	return &Handler1{}
}

func (h *Handler1) SetNext(handler Handler) {
	h.next = handler
}

func (h *Handler1) Next() Handler {
	return h.next
}

func (h *Handler1) HandleRequest(request Request) {
	fmt.Println("handler1..." + request)
	next := h.Next()
	if next != nil {
		next.HandleRequest(request)
	}
}

type Handler2 struct {
	next Handler
}

func NewRequestHandler2() *Handler2 {
	return &Handler2{}
}

func (h *Handler2) SetNext(handler Handler) {
	h.next = handler
}

func (h *Handler2) Next() Handler {
	return h.next
}

func (h *Handler2) HandleRequest(request Request) {
	fmt.Println("handler2..." + request)
	next := h.Next()
	if next != nil {
		next.HandleRequest(request)
	}
}
