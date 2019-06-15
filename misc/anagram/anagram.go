package main

func containSame(sR, tR []rune) bool {
	l := len(sR)
	if l != len(tR) {
		return false
	}
	for i := 0; i < l; i++ {
		if sR[i] != sR[i] {
			return false
		}
	}
	return true
}

func sort(sRary *[]rune) {
	sR := *sRary
	for i := 0; i < len(sR); i++ {
		for j := i; j < len(sR); j++ {
			if sR[i] > sR[j] {
				sR[i], sR[j] = sR[j], sR[i]
			}
		}
	}
}

func areAnagrams(s, t string) bool {
	sR := []rune(s)
	tR := []rune(t)

	sort(&sR)
	sort(&tR)

	return containSame(sR, tR)
}

func main() {
	a := "patrick"
	b := "rickatp"
	c := "rick"

	println(areAnagrams(a, b))
	println(areAnagrams(a, c))
	println(areAnagrams(c, b))
}
