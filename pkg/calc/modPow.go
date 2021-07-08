package calc

func modPow(x, n, mod int) int {
	if n == 0 {
		return 1
	}
	res := modPow(x*x%mod, n/2, mod)
	if n%2 == 1 {
		res = res * x % mod
	}
	return res
}
