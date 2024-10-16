package mwe

import "math/rand"

func Envelopes(quota, count int) (s []int, e error) {
	if 1 == count {
		return []int{quota}, nil
	}

	cur := quota

	for i := 0; i < count-1; i++ {
		n := rand.Intn((2 * cur) / (count - i))
		if n < 1 {
			n = 1 //兜底quota
		}
		s = append(s, n)
		cur -= n
	}
	s = append(s, cur)

	rand.Shuffle(len(s), func(i, j int) {
		s[i], s[j] = s[j], s[i]
	})

	return
}
