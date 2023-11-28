package main

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"
)

/*
 * Complete the 'flippingMatrix' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts 2D_INTEGER_ARRAY matrix as parameter.
 */

func sumSubMatrix(lenSubArr int, matrix [][]int32) int32 {
	var maxTem int32 = 0
	for y := 0; y < lenSubArr; y++ {
		for x := 0; x < lenSubArr; x++ {
			maxTem += matrix[x][y]
		}
	}
	return maxTem
}

func reverseRow(columnNum int, matrix [][]int32) [][]int32 {
	for init, last := 0, len(matrix)-1; init < last; init, last = init+1, last-1 {
		matrix[columnNum][init], matrix[columnNum][last] = matrix[columnNum][last], matrix[columnNum][init]
	}
	return matrix
}

func reverseColumn(rowNum int, matrix [][]int32) [][]int32 {
	for init, last := 0, len(matrix)-1; init < last; init, last = init+1, last-1 {
		matrix[init][rowNum], matrix[last][rowNum] = matrix[last][rowNum], matrix[init][rowNum]
	}
	return matrix
}

func flippingMatrix(matrix [][]int32) int32 {
	// Write your code here
	// 00 01 02 03
	// 10 11 12 13
	// 20 21 22 23
	// 30 31 32 33
	lenArr := len(matrix)
	lenSubArrQuotient := len(matrix) / 2 // 4/2
	var maxSum int32
	for i := 0; i < lenSubArrQuotient; i++ {
		for j := 0; j < lenSubArrQuotient; j++ {
			var temp []int32
			temp = append(temp, matrix[i][j])                   //00 - 01 - 10 - 11
			temp = append(temp, matrix[i][lenArr-j-1])          //03 - 02 - 13 - 12
			temp = append(temp, matrix[lenArr-i-1][j])          //30 - 31 - 20 - 21
			temp = append(temp, matrix[lenArr-i-1][lenArr-j-1]) //33 - 32 - 23 - 22

			sort.Slice(temp, func(i, j int) bool { return temp[i] < temp[j] })

			maxSum += temp[len(temp)-1]
		}
	}
	return maxSum
}

func main() {
	/*
		reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

		stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
		checkError(err)

		defer stdout.Close()

		writer := bufio.NewWriterSize(stdout, 16*1024*1024)

		qTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	*/
	qTemp, err := strconv.ParseInt(strings.TrimSpace("1"), 10, 64)
	checkError(err)
	q := int32(qTemp)

	for qItr := 0; qItr < int(q); qItr++ {
		nTemp, err := strconv.ParseInt(strings.TrimSpace("2"), 10, 64)
		checkError(err)
		n := int32(nTemp)

		var array = [4][1]string{{"112 42 83 119"}, {"56 125 56 49"}, {"15 78 101 43"}, {"62 98 114 108"}}
		var matrix [][]int32
		for i := 0; i < 2*int(n); i++ {
			//matrixRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")
			matrixRowTemp := strings.Split(strings.TrimRight(array[i][0], " \t\r\n"), " ")

			var matrixRow []int32
			for _, matrixRowItem := range matrixRowTemp {
				matrixItemTemp, err := strconv.ParseInt(matrixRowItem, 10, 64)
				checkError(err)
				matrixItem := int32(matrixItemTemp)
				matrixRow = append(matrixRow, matrixItem)
			}

			if len(matrixRow) != 2*int(n) {
				panic("Bad input")
			}

			matrix = append(matrix, matrixRow)
		}

		result := flippingMatrix(matrix)

		fmt.Printf("Result: %d\n", result)
	}

	//writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
