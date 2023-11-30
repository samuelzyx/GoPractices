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
 * Complete the 'caesarCipher' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts following parameters:
 *  1. STRING s
 *  2. INTEGER k
 */

// s clearText & k alphabet rotation factor
func caesarCipher(s string, k int32, n int32) string {
	// Write your code here
	runeArray := []rune(s) //transform to rune []
	var cryptArray []rune

	for _, e := range runeArray {
		//fmt.Println(e)
		tempAscii := int(e) //transform to ASCII number

		// if k is greater that 26 get remainder and use intead of k
		if k > 26 {
			rem := k % 26
			fmt.Println(rem)
			k = rem
		}

		// encript lower cases
		//97 is a & 122 is z
		if tempAscii > 64 && tempAscii < 91 {
			tempAscii += int(k)
			if tempAscii > 90 {
				tempAscii -= 26
			}
		}

		// encript upper cases
		//65 is A & 90 is Z
		if tempAscii > 96 && tempAscii < 123 {
			tempAscii += int(k)
			if tempAscii > 122 {
				tempAscii -= 26
			}
		}
		cryptArray = append(cryptArray, rune(tempAscii))
	}

	return string(cryptArray)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	s := readLine(reader)

	kTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	k := int32(kTemp)

	result := caesarCipher(s, k, n)

	fmt.Fprintf(writer, "%s\n", result)

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
