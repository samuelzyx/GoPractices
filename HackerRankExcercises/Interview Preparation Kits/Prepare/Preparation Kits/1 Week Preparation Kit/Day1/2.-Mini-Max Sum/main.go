package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
 * Complete the 'miniMaxSum' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func miniMaxSum(arr []int32) {
	// Write your code here
	// Sort Slice
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })

	var minSum int64 = 0
	var maxSum int64 = 0

	for index, num := range arr {
		numInt64 := int64(num)
		// Remove minimum number of array to get maximum sum (first index)
		if index != 0 {
			maxSum = maxSum + numInt64
		}
		// Remove maximum number of array to get minimum sum (last index)
		if index != len(arr)-1 {
			minSum = minSum + numInt64
		}
	}

	fmt.Printf("%d %d", minSum, maxSum)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int32

	for i := 0; i < 5; i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	miniMaxSum(arr)
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
