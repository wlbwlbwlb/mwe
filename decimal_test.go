package tool

import "testing"

func TestDecimal_Add(t *testing.T) {
	t.Log(NewDecimal(0.222).Add(0.333).Truncate(2).Float64())
}

func TestDecimal_Sub(t *testing.T) {
	t.Log(NewDecimal(0.222).Sub(0.333).Truncate(2).Float64())
}

func TestDecimal_Mul(t *testing.T) {
	t.Log(NewDecimal(0.222).Mul(3).Truncate(2).Float64())
}
