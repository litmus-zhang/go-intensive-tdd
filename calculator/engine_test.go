package calculator_test

import (
	"log"
	"testing"

	"github.com/litmus-zhang/go-intensive-tdd/calculator"
)

func init() {
	log.Println("Set up")
}

func TestAdd(t *testing.T) {
	e := calculator.Engine{}
	actAssert := func(x, y, expect float64) {
		got := e.Add(x, y)
		if got != expect {
			t.Errorf("Expected %v, got %v", expect, got)
		}

	}
	t.Run("Positive inputs", func(t *testing.T) {
		x, y := 2.5, 3.5
		expect := 6.0
		actAssert(x, y, expect)
	})
	t.Run("Negative inputs", func(t *testing.T) {
		x, y := -2.5, -3.5
		expect := -6.0
		actAssert(x, y, expect)
	})
}

func BenchmarkAdd(b *testing.B) {
	e := calculator.Engine{}

	for i := 0; i < b.N; i++ {
		e.Add(2,3)
	}

}