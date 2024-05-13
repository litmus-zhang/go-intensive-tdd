package arrayslices

import "testing"

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	got := Sum(numbers)
	expect := 15

	if got != expect {
		t.Errorf("got %d expected %d given %v", got, expect, numbers)
	}
}
