package main

import "fmt"

func reverseSection(runeList *[]rune, s, e int) {
	r := *runeList
	for i := 0; i <= (e-s)/2; i++ {
		r[s+i], r[e-i] = r[e-i], r[s+i]
	}
}

func reverse(runeList *[]rune) {
	reverseSection(runeList, 0, len(*runeList)-1)
}

func reversePhrase(runeList *[]rune) {

	r := *runeList
	reverse(runeList)

	start, end := 0, 0
	for i := 0; i < len(r); i++ {
		if r[i] == ' ' {
			end = i - 1
			reverseSection(runeList, start, end)
			start = i + 1
		}
	}
	reverseSection(runeList, start, len(r)-1)

}

func main() {
	// me for sentence a just is this
	phrase := "this is just a sentence for me"

	phraseR := []rune(phrase)
	reversePhrase(&phraseR)

	fmt.Println(string(phraseR))

}
