package mwe

import "testing"

func TestAssert(t *testing.T) {
	var handlers []func()
	Assert(len(handlers) > 0, "must be at least one handler")
}
