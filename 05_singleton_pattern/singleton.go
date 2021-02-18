package main

import (
	"sync"
)

type Counter struct {
	cnt  int
	lock sync.RWMutex
}

func (c *Counter) Add() {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.cnt++
}

func (c *Counter) Count() int {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return c.cnt
}

var (
	once    sync.Once
	counter *Counter
)

func GetCounter() *Counter {
	once.Do(func() {
		counter = &Counter{cnt: 0}
	})
	return counter
}
