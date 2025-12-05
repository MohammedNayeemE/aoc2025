package day5

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type interval struct {
	start, end int
}

func extractFile(filePath string) ([]interval, []int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var ranges []interval
	var ids []int
	readingRanges := true
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			readingRanges = false
			continue
		}
		if readingRanges {
			parts := strings.Split(line, "-")
			start, _ := strconv.Atoi(parts[0])
			end, _ := strconv.Atoi(parts[1])
			ranges = append(ranges, interval{start, end})
		} else {
			id, _ := strconv.Atoi(line)
			ids = append(ids, id)
		}
	}
	return ranges, ids, nil
}

func mergeIntervals(ranges []interval) []interval {
	if len(ranges) == 0 {
		return ranges
	}
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].start < ranges[j].start
	})
	merged := []interval{ranges[0]}
	for _, curr := range ranges[1:] {
		last := &merged[len(merged)-1]
		if curr.start <= last.end+1 {
			if curr.end > last.end {
				last.end = curr.end
			}
		} else {
			merged = append(merged, curr)
		}
	}
	return merged
}

func solve(ranges []interval, ids []int) {
	partOne, partTwo := int64(0), int64(0)
	for _, id := range ids {
		for _, r := range ranges {
			if id >= r.start && id <= r.end {
				partOne++
				break
			}
		}
	}
	for _, r := range ranges {
		partTwo += int64(r.end) - int64(r.start) + 1
	}
	fmt.Println(partOne, partTwo)
}

func Run() {
	ranges, ids, err := extractFile("day5/input.txt")
	if err != nil {
		panic(err)
	}
	ranges = mergeIntervals(ranges)
	solve(ranges, ids)
}
