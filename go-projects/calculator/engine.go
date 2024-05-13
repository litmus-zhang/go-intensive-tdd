package calculator

import "math"

type Engine struct{}

func (e *Engine) Add(x, y float64) float64 {
	return x + y
}

func (e *Engine) Subtract(x, y float64) float64 {
	return math.Abs(x - y)
}

func (e *Engine) Divide(x, y float64) float64 {
	if x == 0 || y == 0 {
		return 0
	}
	return x / y
}

func (e *Engine) Multiplt(x, y float64) float64 {
	return x * y
}
