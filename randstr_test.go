package tool

import "testing"

func TestRandStr(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Log(RandStr(6))
	}
}
