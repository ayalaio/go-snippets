package main

func oddPalindrome(r []rune, idx int) string {

	p := string(r[idx])

	for i := 1; idx-i >= 0 && idx+i < len(r); i++ {
		if r[idx-i] != r[idx+i] {
			break
		}
		p = string(r[idx-i]) + p + string(r[idx+i])
	}

	return p

}

func evenPalindrome(r []rune, idx int) string {

	p := ""

	for i := 0; idx-i >= 0 && idx+i+1 < len(r); i++ {
		if r[idx-i] != r[idx+i+1] {
			break
		}
		p = string(r[idx-i]) + p + string(r[idx+i+1])
	}

	return p

}

func longestPalindrome(s string) string {

	r := []rune(s)

	longest := ""

	for i := 0; i < len(r); i++ {
		p := oddPalindrome(r, i)
		if len(p) > len(longest) {
			longest = p
		}
		q := evenPalindrome(r, i)
		if len(q) > len(longest) {
			longest = q
		}

	}

	return longest

}

func main() {}
