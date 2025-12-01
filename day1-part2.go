package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Movement struct {
	distance int32
	side     string
}

func extractFile() ([]Movement, error) {
	var movements []Movement
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		side := line[:1]
		distance, err := strconv.Atoi(line[1:])
		if err != nil {
			return nil, err
		}
		new_move := Movement{
			distance: int32(distance),
			side:     side,
		}
		movements = append(movements, new_move)
	}
	return movements, nil
}

func calculateNoOfZeros(list []Movement) (int32, error) {
	dial, cnt := int32(50), int32(0)

	for _, move := range list {
		for i := 0; i < int(move.distance); i++ {
			if move.side == "R" {
				dial = (dial + 1) % 100
			} else {
				dial = (dial - 1 + 100) % 100
			}
			if dial == 0 {
				cnt++
			}
		}
	}

	return cnt, nil
}

func main() {
	list, err := extractFile()
	if err != nil {
		fmt.Println(err)
	}
	res, err := calculateNoOfZeros(list)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}
