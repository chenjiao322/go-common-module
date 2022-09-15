package utils

func Set[T comparable](a []T) map[T]struct{} {
	ans := make(map[T]struct{}, 0)
	for _, v := range a {
		ans[v] = struct{}{}
	}
	return ans
}

func Uniq[T comparable](a []T) []T {
	set := Set(a)
	ans := make([]T, 0)
	for k := range set {
		ans = append(ans, k)
	}
	return ans
}

func Diff[T comparable](a, b []T) []T {
	// 在a里但是不在b里
	var ab []T
	sb := Set(b)
	for _, v := range a {
		if _, ok := sb[v]; !ok {
			ab = append(ab, v)
		}
	}
	return ab
}
