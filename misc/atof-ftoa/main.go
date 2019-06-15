package main

import (
	"fmt"
	"math"
)

func atof(a string) float64 {
	f := 0.0
	paddingWhenDotFound := 0
	for i := len(a) - 1; i >= 0; i-- {
		if a[i] == '.' {
			quantOfDecimals := float64(len(a) - 1 - i)
			f = f / math.Pow(10.0, quantOfDecimals)
			paddingWhenDotFound = int(quantOfDecimals) + 1
			continue
		}
		f += (float64(a[i]) - '0') * math.Pow(10.0, float64(len(a)-1-i-paddingWhenDotFound))
	}

	return f
}

func ftoa(f float64, presicionPoint int) string {
	i := ""
	timesDivided := 0
	for f >= 1 {
		timesDivided++
		d := int(f) % 10
		i = string(d+'0') + i
		f = (f - float64(d)) / 10.0
		fmt.Println(i)
	}

	f = f * math.Pow(10.0, float64(timesDivided+presicionPoint))

	if presicionPoint > 0 {
		i = i + "." + ftoa(f, 0)
	}

	fmt.Println(i)

	return i

}

func main() {
	a := "123.456"
	f := 123.456

	res_atof := atof(a)
	res_ftoa := ftoa(f, 3)

	if res_atof != f {
		fmt.Println("Error on atof")
	}

	if res_ftoa != a {
		fmt.Println("Error on ftoa")
	}

}
