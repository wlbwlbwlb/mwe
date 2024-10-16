package mwe

import "testing"

func TestPages(t *testing.T) {
	t.Log(Pages(10, 0))
	t.Log(Pages(10, 1))
	t.Log(Pages(10, 9))
	t.Log(Pages(10, 10))
	t.Log(Pages(10, 11))
}
