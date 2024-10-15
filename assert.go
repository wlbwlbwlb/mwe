package util

func Assert(expr bool, tip string) {
	if !expr {
		panic(tip)
	}
}
