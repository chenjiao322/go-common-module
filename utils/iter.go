package utils

func Map[T1 any, T2 any](f func(T1) T2, slice []T1) []T2 {
	ans := make([]T2, len(slice))
	for i, v := range slice {
		ans[i] = f(v)
	}
	return ans
}

func Filter[T any](f func(T) bool, slice []T) []T {
	ans := make([]T, 0)
	for _, v := range slice {
		if f(v) {
			ans = append(ans, v)
		}
	}
	return ans
}

func Fold[T any, Acc any](f func(T, Acc) Acc, a Acc, slice []T) Acc {
	for _, v := range slice {
		a = f(v, a)
	}
	return a
}

func ForEach[T any](f func(T), slice []T) {
	for _, v := range slice {
		f(v)
	}
}

func Zip[T any](slices ...[]T) [][]T {
	size := Min(Map(func(a []T) int { return len(a) }, slices)...)
	ans := make([][]T, size)
	for i := 0; i < size; i++ {
		ans[i] = Map(func(s []T) T { return s[i] }, slices)
	}
	return ans
}

func Reverse[T any](slice []T) {
	n := len(slice)
	for i := 0; i < n/2; i++ {
		slice[i], slice[n-i-1] = slice[n-i-1], slice[i]
	}
}

func GroupBy[T any, cp comparable](f func(T) cp, slice []T) [][]T {
	ans := make([][]T, 0)
	if len(slice) == 0 {
		return ans
	}
	pre := f(slice[0])
	tmp := []T{slice[0]}
	for i := 1; i < len(slice); i++ {
		cur := f(slice[i])
		if pre == cur {
			tmp = append(tmp, slice[i])
		} else {
			ans = append(ans, tmp)
			tmp = []T{slice[i]}
		}
		pre = cur
	}
	if len(tmp) > 0 {
		ans = append(ans, tmp)
	}
	return ans
}
