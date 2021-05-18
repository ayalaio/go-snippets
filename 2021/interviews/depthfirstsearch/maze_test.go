package depthfirstsearch

import (
	"errors"
	"testing"
)

type position struct {
	x int
	y int
}

func SolveMazeShortestPath(maze [][]int, visited [][]bool, start, end position) ([]position, error) {
	err := errors.New("Imposible to solve")
	if start.x >= len(maze) ||
		start.x < 0 ||
		start.y >= len(maze[0]) ||
		start.y < 0 {
		return nil, err
	}
	if end.x >= len(maze) ||
		end.x < 0 ||
		end.y >= len(maze[0]) ||
		end.y < 0 {
		return nil, err
	}

	if visited[start.x][start.y] {
		return nil, err
	}
	visited[start.x][start.y] = true

	if maze[start.x][start.y] == 1 {
		return nil, err
	}
	if maze[end.x][end.y] == 1 {
		return nil, err
	}

	currentPosition := position{start.x, start.y}
	if start.x == end.x && start.y == end.y {
		return []position{currentPosition}, nil
	}

	sol1, err1 := SolveMazeShortestPath(maze, visited, position{start.x - 1, start.y}, end)
	sol2, err2 := SolveMazeShortestPath(maze, visited, position{start.x + 1, start.y}, end)
	sol3, err3 := SolveMazeShortestPath(maze, visited, position{start.x, start.y - 1}, end)
	sol4, err4 := SolveMazeShortestPath(maze, visited, position{start.x, start.y + 1}, end)

	var sol []position = nil
	if err1 == nil && (len(sol1) < len(sol) || sol == nil) {
		sol = sol1
	}
	if err2 == nil && (len(sol2) < len(sol) || sol == nil) {
		sol = sol2
	}
	if err3 == nil && (len(sol3) < len(sol) || sol == nil) {
		sol = sol3
	}
	if err4 == nil && (len(sol4) < len(sol) || sol == nil) {
		sol = sol4
	}

	if sol != nil {
		return append([]position{currentPosition}, sol...), nil
	}
	return nil, err
}

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
