package mwe

import "testing"

func TestSum(t *testing.T) {
	t.Log(Sum([]int{1, 2, 3}))
}

func TestIn(t *testing.T) {
	t.Log(In(2, []int{1, 2, 3}))
}
