package day7

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func extractFile(filePath string) ([][]rune, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var grid [][]rune
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		grid = append(grid, []rune(line))

	}
	return grid, nil
}

func sum(nums []int) int {
	s := 0
	for _, num := range nums {
		s += num
	}
	return s
}

func solve(grid [][]rune) {
	m, n := len(grid), len(grid[0])
	vi := make([]int, n)
	vj := make([]int, n)
	partOne, partTwo := 0, 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 'S' {
				vi[j] = 1
			}
			if grid[i][j] == '^' {
				if vi[j] == 0 {
					continue
				}
				done := false
				if j-1 >= 0 {
					vi[j-1] = 1
					done = true
				}
				if j+1 < n {
					vi[j+1] = 1
					done = true
				}
				if done {
					vi[j] = 0
					partOne++
				}
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 'S' {
				vj[j] = 1
			}
			if grid[i][j] == '^' {
				if vj[j] == 0 {
					continue
				}
				if j-1 >= 0 {
					vj[j-1] += vj[j]
				}
				if j+1 < n {
					vj[j+1] += vj[j]
				}
				vj[j] = 0
			}
		}
	}

	partTwo += sum(vj)
	fmt.Println(partOne, partTwo)
}

func Run() {
	grid, err := extractFile("day7/input.txt")
	if err != nil {
		panic(err)
	}
	solve(grid)
}
