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
 * Complete the 'timeConversion' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func timeConversion(s string) string {
	// Write your code here
	// Check if is PM or AM
	isPM := strings.Contains(s, "PM")
	isAM := strings.Contains(s, "AM")
	getHour, _ := strconv.Atoi(s[0:2])
	getRest := s[2:8] // substring [:mm:ss]
	// If is PM
	if isPM {
		if getHour == 12 {
			// If is PM 12:00:00 it will be 12:00:00
			return "12" + getRest // concat [:mm:ss]
		} else {
			// Add to 12 to hours
			getHour += 12
			// If is AM or PM remove (AM/PM) substring
			return strconv.Itoa(getHour) + getRest // concat [:mm:ss]
		}
	}
	if isAM {
		// If is AM 12:00:00 it will be 00:00:00
		if getHour == 12 {
			// If is PM 12:00:00 it will be 12:00:00
			return "00" + getRest // concat [:mm:ss]
		}
	}
	// If is AM or PM remove (AM/PM) substring
	return "0" + strconv.Itoa(getHour) + getRest // concat [:mm:ss]
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	s := readLine(reader)

	result := timeConversion(s)

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
