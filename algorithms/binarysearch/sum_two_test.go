package binarysearch

import "testing"

func TestSumTwoImpl(t *testing.T) {
	cases := []struct {
		numbers []int
		sum     int
		result  []int
	}{
		{
			numbers: []int{1, 2, 3, 4, 5},
			sum:     9,
			result:  []int{3, 4},
		},
		{
			numbers: []int{1, 2, 3, 5},
			sum:     7,
			result:  []int{1, 3},
		},
	}
	for _, c := range cases {
		left, right, err := sumTwoImpl(c.numbers, c.sum)
		if err != nil {
			t.Errorf("got %v", err)
		}
		if left != c.result[0] || right != c.result[1] {
			t.Errorf("got %d %d", left, right)
			t.Errorf("expected %d %d", c.result[0], c.result[1])
		}
	}
}
