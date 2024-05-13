package generics

import "testing"

func TestAssertFunctions(t *testing.T) {
	t.Run("Asserting integers", func(t *testing.T) {
		AssertEqual(t, 1, 1)
		AssertNotEqual(t, 1, 2)
	})
	t.Run("Asserting strings", func(t *testing.T) {
		AssertEqual(t, "hello", "hello")
		AssertNotEqual(t, "hello", "world")
		AssertEqual(t, 1, 1)
	})
}

func AssertEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func AssertNotEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got == want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func AssertTrue(t *testing.T, got bool) {
	t.Helper()
	if !got {
		t.Errorf("expected true but got false")
	}
}

func AssertFalse(t *testing.T, got bool) {
	t.Helper()
	if got {
		t.Errorf("expected false but got true")
	}
}

func TestStack(t *testing.T) {
	t.Run("Stack of integers", func(t *testing.T) {
		myStackOfInts := new(Stack[int])

		AssertTrue(t, myStackOfInts.IsEmpty())

		myStackOfInts.Push(123)
		myStackOfInts.Push(456)
		AssertFalse(t, myStackOfInts.IsEmpty())
		value, _ := myStackOfInts.Pop()
		AssertEqual(t, value, 456)
		value, _ = myStackOfInts.Pop()
		AssertEqual(t, value, 123)
		AssertTrue(t, myStackOfInts.IsEmpty())
	})
	t.Run("Interface stack DX is horrid", func(t *testing.T) {
		myStackOfInts := new(Stack[int])
		myStackOfInts.Push(1)
		myStackOfInts.Push(2)
		first, _ := myStackOfInts.Pop()
		second, _ := myStackOfInts.Pop()
		AssertEqual(t, first+second, 3)
	})
}
