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
 * Complete the 'fizzBuzz' function below.
 *
 * The function accepts INTEGER n as parameter.
 */

func fizzBuzz(n int32) {
	// Write your code here
	for i := 1; i <= int(n); i++ {
		var s string = ""
		var isMult bool = false
		if i%3 == 0 {
			s += "Fizz"
			isMult = true
		}
		if i%5 == 0 {
			s += "Buzz"
			isMult = true
		}
		if !isMult {
			s += fmt.Sprintf("%d", i)
		}
		fmt.Println(s)
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	fizzBuzz(n)
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
