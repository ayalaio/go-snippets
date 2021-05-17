package maze

import (
	"testing"
)

func TestMazeReturnsShortestPathWhenSuccessful(t *testing.T) {
	maze := [][]int{
		{0, 1, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 0, 1, 0},
		{0, 1, 0, 0, 0, 1, 0},
		{0, 1, 0, 0, 0, 1, 0},
		{0, 1, 1, 1, 0, 1, 0},
		{0, 0, 0, 0, 0, 1, 0},
		{0, 1, 1, 1, 1, 1, 0},
	}
	visited := make([][]bool, 7)
	for i := range visited {
		visited[i] = make([]bool, 7)
	}
	sol, _ := SolveMazeShortestPath(maze, visited, position{0, 0}, position{6, 6})
	if len(sol) != 23 {
		t.Errorf("Expected 23 to be the length of the solution")
	}
}
