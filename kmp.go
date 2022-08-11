package kmp

func getNext[T comparable](p []T) []int {
	l := len(p)
	j := 0
	k := -1
	m := l - 1
	next := make([]int, l, l)
	next[0] = -1
	for j < m {
		if k == -1 || p[j] == p[k] {
			j++
			k++
			if p[j] == p[k] {
				next[j] = next[k]
			} else {
				next[j] = k
			}
		} else {
			k = next[k]
		}
	}
	return next
}

func Index[T comparable](test, pattern []T) int {
	lt := len(test)
	lp := len(pattern)
	if lt < lp || lp == 0 {
		return -1
	}
	i := 0
	j := 0
	next := getNext(pattern)
	for i < lt && j < lp {
		if j == -1 || test[i] == pattern[j] {
			i++
			j++
		} else {
			j = next[j]
		}
	}
	if j == lp {
		return i - j
	}
	return -1
}

func Indexes[T comparable](test, pattern []T) []int {
	var res []int
	lp := len(pattern)
	i := Index(test, pattern)
	j := 0
	for i != -1 {
		j += i
		res = append(res, j)
		j += lp
		i = Index(test[j:], pattern)
	}
	return res
}
