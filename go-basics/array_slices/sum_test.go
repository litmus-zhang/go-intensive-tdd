package arrayslices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	got := Sum(numbers)
	expect := 15

	if got != expect {
		t.Errorf("got %d expected %d given %v", got, expect, numbers)
	}
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	expect := []int{3, 9}
	if !reflect.DeepEqual(got, expect) {
		t.Errorf("got %v expected %v", got, expect)
	}
}

func TestSumAllTails(t *testing.T) {
	got := SumAllTail([]int{1, 2}, []int{0, 9})
	expect := []int{2, 9}

	if !reflect.DeepEqual(got, expect) {
		t.Errorf("got %v expected %v", got, expect)
	}
}
