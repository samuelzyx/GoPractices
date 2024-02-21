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
 * Complete the 'superDigit' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. STRING n
 *  2. INTEGER k
 */

func reduce(n string) string {
	result := 0
	charNumber := []rune(n)
	for i := 0; i < len(charNumber); i++ {
		val, _ := strconv.Atoi(string(charNumber[i]))
		result += val
	}
	fmt.Printf("n: %s\n", n)
	fmt.Printf("result: %d\n", result)
	if result > 10 {
		return reduce(fmt.Sprintf("%d", result))
	} else {
		return fmt.Sprintf("%d", result)
	}
}

func superDigit(n string, k int32) int32 {
	// Write your code here
	fmt.Printf("k: %d\n", k)
	finalS := ""
	for i := 0; i < int(k); i++ {
		finalS += n
	}
	finalS = reduce(finalS)
	val, _ := strconv.Atoi(finalS)
	return int32(val)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	n := firstMultipleInput[0]

	kTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	k := int32(kTemp)

	result := superDigit(n, k)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
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
