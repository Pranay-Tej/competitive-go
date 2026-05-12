package bitwise

import "testing"

func TestCaseOne(t *testing.T) {
	PowerSet([]int{4, 7})
}

func TestCaseTwo(t *testing.T) {
	PowerSet([]int{-1, 0, 1, 2})
}
