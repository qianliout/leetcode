package main

func main() {

}

func superPow(a int, b []int) int {
	mod := 1337
	if len(b) == 0 {
		return 1
	}
	if len(b) == 1 {
		return pow(a, b[0], mod)
	}
	last := b[len(b)-1]
	b = b[:len(b)-1]
	return (pow(superPow(a, b), 10, mod) * pow(a, last, mod)) % mod
}

func pow(a, b, base int) int {
	if b == 0 {
		return 1
	}

	a = a % base

	if b%2 == 1 {
		return (a * pow(a, b-1, base)) % base
	} else {
		sub := pow(a, b>>1, base)
		return (sub * sub) % base
	}
}
