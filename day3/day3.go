package day3

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func extractFile(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	var list []string
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		list = append(list, line)
	}
	return list, nil
}

func maxDigit(s string) int {
	max := byte('0')
	maxI := -1
	for i := 0; i < len(s); i++ {
		if s[i] > max {
			max = s[i]
			maxI = i
		}
	}
	return maxI
}

func calculateAnsOne(list []string) {
	res := 0
	for _, num := range list {
		index, largeNum := maxDigit(num), "0"
		rightSide, leftSide := -1, -1
		if index != len(num)-1 {
			rightSide = maxDigit(num[index+1:])
		}
		if index != 0 {
			leftSide = maxDigit(num[:index])
		}
		if rightSide != -1 {
			largeNum = string([]byte{num[index], num[index+1+rightSide]})
		} else if leftSide != -1 {
			largeNum = string([]byte{num[leftSide], num[index]})
		}
		val, _ := strconv.Atoi(largeNum)
		res += val
	}
	fmt.Println(res)
}

var dp map[[2]int]string

func solve(idx int, k int, s string) string {
	if k == 0 {
		return ""
	}
	if len(s[idx:]) == k {
		return s[idx:]
	}
	key := [2]int{idx, k}
	if val, ok := dp[key]; ok {
		return val
	}
	take := string(s[idx]) + solve(idx+1, k-1, s)
	notake := solve(idx+1, k, s)
	dp[key] = max(take, notake)
	return dp[key]
}

func calculateAnsTwo(list []string) {
	res := int64(0)
	for _, num := range list {
		dp = make(map[[2]int]string)
		best := solve(0, 12, num)
		val, _ := strconv.ParseInt(best, 10, 64)
		res += val
	}
	fmt.Println(res)
}

func Run() {
	list, err := extractFile("day3/input.txt")
	if err != nil {
		fmt.Println(err)
	}
	calculateAnsOne(list)
	calculateAnsTwo(list)
}
