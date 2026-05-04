package main

import "fmt"

// #200 Number of Islands [Medium]
// Tags: Array, Depth-First Search, Breadth-First Search, Union Find, Matrix

func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	type Node struct {
		X int
		Y int
	}

	rows := len(grid)
	cols := len(grid[0])
	visited := make(map[Node]bool)
	islands := 0

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '1' && !visited[Node{X: i, Y: j}] {
				stack := []Node{{X: i, Y: j}}
				for len(stack) > 0 {
					curr := stack[len(stack)-1]
					stack = stack[:len(stack)-1]

					if visited[curr] {
						continue
					}

					visited[curr] = true

					// 檢查上下左右
					directions := []Node{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
					for _, dir := range directions {
						ni, nj := curr.X+dir.X, curr.Y+dir.Y
						if ni >= 0 && ni < rows && nj >= 0 && nj < cols &&
							grid[ni][nj] == '1' && !visited[Node{X: ni, Y: nj}] {
							stack = append(stack, Node{X: ni, Y: nj})
						}
					}
				}
				islands++
			}
		}
	}
	return islands
}

func main() {
	fmt.Println(numIslands([][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	})) // → 1

	fmt.Println(numIslands([][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	})) // → 3

	fmt.Println(numIslands([][]byte{
		{'1'},
	})) // → 1

	fmt.Println(numIslands([][]byte{
		{'0'},
	})) // → 0
}
