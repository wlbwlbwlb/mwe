package tool

func Assert(expr bool, tip string) {
	if !expr {
		panic(tip)
	}
}
