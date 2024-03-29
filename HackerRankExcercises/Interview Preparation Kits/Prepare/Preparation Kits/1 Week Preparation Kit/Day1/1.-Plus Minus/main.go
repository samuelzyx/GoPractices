package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'plusMinus' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func plusMinus(arr []int32) {
	// Write your code here
	// Count how many positive number we have in the array
	positiveNum := 0
	// Count how many negative number we have in the array
	negativeNum := 0
	// Count how many zeros we have in the array
	zeroNum := 0
	for _, num := range arr {
		if num == 0 {
			zeroNum++
		}
		if num < 0 {
			negativeNum++
		}
		if num > 0 {
			positiveNum++
		}
	}
	// According the N size of array
	lenArr := len(arr)
	// Calculate ratio (total positive number/n) and print with 6 places after decimal
	fmt.Printf("%.6f\n", float32(positiveNum)/float32(lenArr))
	// Calculate ratio (total negative number/n) and print with 6 places after decimal
	fmt.Printf("%.6f\n", float32(negativeNum)/float32(lenArr))
	// Calculate ratio (total zero/n) and print with 6 places after decimal
	fmt.Printf("%.6f\n", float32(zeroNum)/float32(lenArr))
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	plusMinus(arr)
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
