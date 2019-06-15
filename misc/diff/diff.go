package main

import "math"

func lookFor(c rune, s []rune, pos int) int {
	for i := pos; i < len(s); i++ {
		if s[i] == c {
			return i
		}
	}
	return math.MaxInt64
}

func subseq(s []rune, start, end int) []rune {
	return s[start:end]
}

func diff(_s, _t string) string {

	res := ""

	s := []rune(_s)
	t := []rune(_t)

	i := 0
	j := 0

	for true {

		if j >= len(t) || i >= len(s) {
			for j < len(t) {
				res = res + " +" + string(t[j])
				j++
			}
			for i < len(s) {
				res = res + " -" + string(s[i])
				i++
			}
			break
		}

		println(res)

		if s[i] == t[j] {
			res = res + " " + string(s[i])
			i++
			j++
		} else {
			// look for t[j] in s, gather all sequence
			l := lookFor(t[j], s, i)

			// look for s[i] in t, gather all sequence
			m := lookFor(s[i], t, j)

			if m == math.MaxInt64 && l == math.MaxInt64 {
				i++
				j++
				println(string(s[i]), string(t[j]))
				//res = res + string(s[i])
				continue
			}

			if l < m {
				// add - to seq if found on t, move i + len(subseq)
				for _, c := range subseq(s, i, l) {
					res = res + " -" + string(c)
				}
				i += len(subseq(s, i, l))

			} else {
				// add + to seq if found on s, move j + len(subseq)
				for _, c := range subseq(t, j, m) {
					res = res + " +" + string(c)
				}
				j += len(subseq(t, j, m))
			}
		}
	}

	return res
}

func main() {
	a := "XMJYAUZ"
	b := "XMJAATZ"

	r := diff(a, b)

	println(r)
}
