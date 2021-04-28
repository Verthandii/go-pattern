package main

type NumberStrategy interface {
	Strategy(int, int) int
}

type SumStrategy struct{}

func (st *SumStrategy) Strategy(x, y int) int {
	return x + y
}

type MinusStrategy struct{}

func (mt *MinusStrategy) Strategy(x, y int) int {
	return x - y
}

type Machine struct {
	strategy NumberStrategy
}

func NewMachine(strategy NumberStrategy) *Machine {
	return &Machine{strategy: strategy}
}

func (m *Machine) Calc(x, y int) int {
	return m.strategy.Strategy(x, y)
}
