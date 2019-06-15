package main

import "fmt"

type Coordinate struct {
	x, y int
}

func longestPathSize(a [][]int, src, dest Coordinate) (int, error) {

	if src.x < 0 || src.x >= len(a) || src.y < 0 || src.y >= len(a[0]) || a[src.x][src.y] == 0 {
		return 0, fmt.Errorf("Imposible")
	}

	if src.x == dest.x && src.y == dest.y {
		return 0, nil
	}

	// Make coordinate unwalkable
	a[src.x][src.y] = 0

	maxAry := make([]int, 0)
	for _, el := range []Coordinate{
		Coordinate{src.x - 1, src.y},
		Coordinate{src.x + 1, src.y},
		Coordinate{src.x, src.y - 1},
		Coordinate{src.x, src.y + 1},
	} {
		if res, err := longestPathSize(a, el, dest); err == nil {
			maxAry = append(maxAry, res)
		}
	}

	// Make coordinate walkable again
	a[src.x][src.y] = 1

	max := 0
	err := fmt.Errorf("Imposible")
	for _, el := range maxAry {
		if el >= max {
			max = el
			err = nil
		}
	}

	if err != nil {
		return 0, err
	}

	return max + 1, nil

}

func main() {
	a := [][]int{
		{1, 0, 1, 1, 1, 1, 0, 1, 1, 1},
		{1, 0, 1, 0, 1, 1, 1, 0, 1, 1},
		{1, 1, 1, 0, 1, 1, 0, 1, 0, 1},
		{0, 0, 0, 0, 1, 0, 0, 1, 0, 0},
		{1, 0, 0, 0, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 0},
		{1, 0, 0, 0, 1, 0, 0, 1, 0, 1},
		{1, 0, 1, 1, 1, 1, 0, 0, 1, 1},
		{1, 1, 0, 0, 1, 0, 0, 0, 0, 1},
		{1, 0, 1, 1, 1, 1, 0, 1, 0, 0},
	}

	var lps int
	var err error

	lps, err = longestPathSize(a, Coordinate{0, 0}, Coordinate{5, 7})
	fmt.Println(lps, err)

	lps, err = longestPathSize(a, Coordinate{0, 0}, Coordinate{1, 0})
	fmt.Println(lps, err)

	lps, err = longestPathSize(a, Coordinate{0, 0}, Coordinate{0, 1})
	fmt.Println(lps, err)

}
