package day2

//
// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"strings"
// )
//
// type Range struct {
// 	Start int
// 	End   int
// }
//
// func extractFile(filePath string) ([]Range, error) {
// 	file, err := os.Open(filePath)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer file.Close()
// 	var list []Range
// 	scanner := bufio.NewScanner(file)
// 	for scanner.Scan() {
// 		line := strings.TrimSpace(scanner.Text())
// 		splittedRanges := strings.Split(line, ",")
// 		for _, itr := range splittedRanges {
// 			temp := strings.Split(itr, "-")
// 			if len(temp) != 2 {
// 				continue
// 			}
// 			start, _ := strconv.Atoi(temp[0])
// 			end, _ := strconv.Atoi(temp[1])
// 			list = append(list, Range{
// 				Start: start,
// 				End:   end,
// 			})
// 		}
// 	}
// 	if err := scanner.Err(); err != nil {
// 		return nil, err
// 	}
// 	return list, nil
// }
//
// func calculateAns(list []Range) {
// 	var res int64 = 0
// 	for _, item := range list {
// 		for i := item.Start; i <= item.End; i++ {
// 			s := strconv.Itoa(i)
// 			firstPart, secondPart := s[:len(s)/2], s[len(s)/2:]
// 			if firstPart == secondPart {
// 				res += int64(i)
// 			}
// 		}
// 	}
// 	fmt.Println(res)
// }
//
// func Run(filePath string) {
// 	list, err := extractFile(filePath)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	calculateAns(list)
// }
