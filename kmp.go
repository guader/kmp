package kmp

func getNext[T comparable](pattern []T) []int {
	l := len(pattern)
	j := 0
	k := -1
	m := l - 1
	next := make([]int, l, l)
	next[0] = -1
	for j < m {
		if k == -1 || pattern[j] == pattern[k] {
			j++
			k++
			if pattern[j] == pattern[k] {
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

func index[T comparable](test, pattern []T, next []int) int {
	lt := len(test)
	lp := len(pattern)
	if lt < lp || lp == 0 {
		return -1
	}
	i := 0
	j := 0
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

// Index find index of pattern from test
func Index[T comparable](test, pattern []T) int {
	return index(test, pattern, getNext(pattern))
}

// Indexes find indexes of patterns not overlapped from test
func Indexes[T comparable](test, pattern []T) []int {
	var res []int
	next := getNext(pattern)
	lp := len(pattern)
	i := index(test, pattern, next)
	j := 0
	for i != -1 {
		j += i
		res = append(res, j)
		j += lp
		i = index(test[j:], pattern, next)
	}
	return res
}
