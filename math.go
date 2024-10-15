package tool

type Number interface {
	int64 | int
}

func Sum[T Number](s []T) (sum T) {
	for _, v := range s {
		sum += v
	}
	return
}

func In[T Number](val T, s []T) (got bool) {
	for _, v := range s {
		if val == v {
			return true
		}
	}
	return
}
