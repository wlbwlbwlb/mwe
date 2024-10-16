package mwe

import "testing"

func TestIf(t *testing.T) {
	age := 0
	age = If(age > 0, age, 23)
	t.Log(age)
}
