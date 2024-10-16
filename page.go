package mwe

func Pages(size int, total int64) (n int) {
	i32 := int(total)

	x := i32 / size

	y := i32 % size

	//整页的
	n = x

	//余下的
	if y > 0 {
		n++
	}

	return
}
