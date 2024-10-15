package tool

import "testing"

func TestEnvelopes(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Log(Envelopes(7, 6))
	}
}
