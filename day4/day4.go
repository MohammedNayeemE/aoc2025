package day4

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cell struct {
	r, c int
}

func extractFile(filePath string) ([][]rune, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	var grid [][]rune
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		grid = append(grid, []rune(line))
	}
	return grid, nil
}

func solve(grid [][]rune) {
	partOne, partTwo := 0, 0
	m, n := len(grid), len(grid[0])
	dirs := [8][2]int{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, -1},
		{0, 1},
		{1, -1},
		{1, 0},
		{1, 1},
	}
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] != '@' {
				continue
			}
			neigh := 0
			for _, it := range dirs {
				di, dj := i+it[0], j+it[1]
				if di >= 0 && di < m && dj >= 0 && dj < n && grid[di][dj] == '@' {
					neigh++
				}
			}
			dp[i][j] = neigh
		}
	}
	q := []cell{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '@' && dp[i][j] < 4 {
				q = append(q, cell{r: i, c: j})
			}
		}
	}
	partOne = len(q)
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		r, c := cur.r, cur.c
		if grid[r][c] != '@' {
			continue
		}
		grid[r][c] = '.'
		partTwo++
		for _, it := range dirs {
			di, dj := r+it[0], c+it[1]
			if di >= 0 && di < m && dj >= 0 && dj < n && grid[di][dj] == '@' {
				dp[di][dj]--
				if dp[di][dj] == 3 {
					q = append(q, cell{r: di, c: dj})
				}
			}
		}
	}
	fmt.Println(partOne)
	fmt.Println(partTwo)
}

func Run() {
	grid, err := extractFile("day4/input.txt")
	if err != nil {
		fmt.Println(err)
	}
	solve(grid)
}
