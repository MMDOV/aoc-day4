package main

import (
	"bufio"
	"fmt"
	"os"
)

var xmasCount int
var masCount int

func readFile(filePath string) [][]string {
	var allRows [][]string
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		return allRows
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str := scanner.Text()
		var row []string
		for _, char := range []rune(str) {
			row = append(row, string(char))
		}
		allRows = append(allRows, row)
	}
	return allRows
}

func sliceSubtract(array1, array2 []int) []int {
	if len(array1) != len(array2) {
		panic("Slices must have the same length")
	}

	result := make([]int, len(array1))
	for i := range array1 {
		result[i] = array2[i] - array1[i]
	}
	return result
}

func sliceAdd(array1, array2 []int) []int {
	if len(array1) != len(array2) {
		panic("Slices must have the same length")
	}

	result := make([]int, len(array1))
	for i := range array1 {
		result[i] = array2[i] + array1[i]
	}
	return result
}

func isValid(rows [][]string, pos [2]int) bool {
	return pos[0] >= 0 && pos[0] < len(rows) &&
		pos[1] >= 0 && pos[1] < len(rows[pos[0]])
}

func findXmas(rows [][]string, currPos [2]int, nextPos [2]int) {
	diff := sliceSubtract(currPos[:], nextPos[:])
	for isValid(rows, nextPos) {
		if rows[nextPos[0]][nextPos[1]] == "M" {
			nextPos = [2]int(sliceAdd(diff, nextPos[:]))
			if isValid(rows, nextPos) && rows[nextPos[0]][nextPos[1]] == "A" {
				nextPos = [2]int(sliceAdd(diff, nextPos[:]))
				if isValid(rows, nextPos) && rows[nextPos[0]][nextPos[1]] == "S" {
					xmasCount += 1
				} else {
					return
				}
			} else {
				return
			}
		} else {
			return
		}
	}
}

func findMas(rows [][]string, currPos [2]int) {
	topRight := sliceAdd(currPos[:], []int{1, 1})
	topLeft := sliceAdd(currPos[:], []int{-1, 1})
	buttomRight := sliceAdd(currPos[:], []int{1, -1})
	buttomLeft := sliceAdd(currPos[:], []int{-1, -1})
	if isValid(rows, [2]int(topLeft)) && isValid(rows, [2]int(topRight)) &&
		isValid(rows, [2]int(buttomLeft)) && isValid(rows, [2]int(buttomRight)) {
		if (rows[topRight[0]][topRight[1]] == "S" && rows[buttomLeft[0]][buttomLeft[1]] == "M") ||
			(rows[topRight[0]][topRight[1]] == "M" && rows[buttomLeft[0]][buttomLeft[1]] == "S") {
			if (rows[topLeft[0]][topLeft[1]] == "S" && rows[buttomRight[0]][buttomRight[1]] == "M") ||
				(rows[topLeft[0]][topLeft[1]] == "M" && rows[buttomRight[0]][buttomRight[1]] == "S") {
				masCount += 1
			}
		}
	}
}

func main() {
	allRows := readFile("day4.input")
	for i := 0; i < len(allRows); i++ {
		for j := 0; j < len(allRows[i]); j++ {
			if allRows[i][j] == "X" {
				findXmas(allRows, [2]int{i, j}, [2]int{i + 1, j})
				findXmas(allRows, [2]int{i, j}, [2]int{i - 1, j})
				findXmas(allRows, [2]int{i, j}, [2]int{i, j + 1})
				findXmas(allRows, [2]int{i, j}, [2]int{i, j - 1})
				findXmas(allRows, [2]int{i, j}, [2]int{i + 1, j + 1})
				findXmas(allRows, [2]int{i, j}, [2]int{i - 1, j + 1})
				findXmas(allRows, [2]int{i, j}, [2]int{i + 1, j - 1})
				findXmas(allRows, [2]int{i, j}, [2]int{i - 1, j - 1})
			}
			if allRows[i][j] == "A" {
				findMas(allRows, [2]int{i, j})
			}
		}
	}
	fmt.Println("xmas count:")
	fmt.Println(xmasCount)
	fmt.Println("mas count:")
	fmt.Println(masCount)
}
