package mwe

func If[T any](expr bool, t T, f T) T {
	if expr {
		return t
	}
	return f
}
