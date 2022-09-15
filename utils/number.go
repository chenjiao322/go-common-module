package utils

func IsInt(s string) bool {
	if len(s) == 0 {
		return false
	}
	for _, v := range s {
		if v < '0' || '9' < v {
			return false
		}
	}
	return true
}

func Max[T Number](a ...T) T {
	if len(a) == 0 {
		panic("Max expected more than 1 argument,got 0")
	}
	ma := a[0]
	for _, v := range a {
		if v > ma {
			ma = v
		}
	}
	return ma
}

func Min[T Number](a ...T) T {
	if len(a) == 0 {
		panic("Min expected more than 1 argument,got 0")
	}
	mi := a[0]
	for _, v := range a {
		if v < mi {
			mi = v
		}
	}
	return mi
}

func Sum[T Number](a ...T) T {
	return Fold(func(t T, acc T) T { return t + acc }, T(0), a)
}

func Gcd(a, b uint64) uint64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func Pow(base, exp uint64, mod uint32) uint64 {
	MOD := uint64(mod) // mod 必须小于uint32才能保证计算过程中不溢出
	if exp == 0 {
		return 1
	}
	base %= MOD
	ans := base
	exp -= 1
	for exp > 0 {
		if exp&1 > 0 {
			ans = (base * ans) % MOD
		}
		base = (base * base) % MOD
		exp >>= 1
	}
	return ans
}
