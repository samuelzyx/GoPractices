package csvreader

import (
	"bufio" // Used for reading user input line by line
	"flag"
	"fmt"     // Used for printing messages and formatting output
	"os"      // Used for file operations and reading the CSV file
	"strconv" // Used for converting user input (string) to a number (int) for answer checking
	"strings"
	"time" // Used for handling timers
)

type question struct {
	text   string // Text of the question
	answer string // Correct answer to the question
}

// Reads a CSV file containing quiz questions and returns a slice of question structs
func readCSV(fileName string) ([]question, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var questions []question
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		if len(parts) != 2 {
			continue // Ignore invalid lines
		}

		q := question{
			text:   parts[0],
			answer: strings.TrimSpace(parts[1]),
		}
		questions = append(questions, q)
	}

	return questions, nil
}

// Asks a question to the user, reads their answer, and checks if it's correct
func askQuestion(i int, q question) bool {
	fmt.Printf("Problem #%d: %s=\n", i, q.text)
	var userAnswer string
	fmt.Scanln(&userAnswer)

	trimmedAnswer := strings.TrimSpace(userAnswer)
	correctAnswer, err := strconv.Atoi(q.answer)
	if err != nil {
		return false // Invalid answer format
	}

	userAnswerInt, err := strconv.Atoi(trimmedAnswer)
	if err != nil {
		return false // Invalid answer format
	}

	return userAnswerInt == correctAnswer
}

// Creates a channel that sends a signal after a specified duration (used for timeout)
func timeout(d time.Duration) <-chan time.Time {
	return time.After(d)
}

func main() {
	fileName := "problems.csv" // Default filename for the CSV quiz questions

	// Flag to allow specifying a custom CSV file name
	flag.StringVar(&fileName, "file", fileName, "CSV file containing quiz questions")

	var duration time.Duration = 30 * time.Second // Default timeout duration (30 seconds)

	// Flag to allow specifying a custom time limit
	flag.DurationVar(&duration, "time", duration, "Time limit for the quiz (seconds)")
	flag.Parse()

	timeoutCh := timeout(duration)

	fmt.Println("Press Enter to start the quiz...")
	_, err := fmt.Scanln()
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	questions, err := readCSV(fileName)
	if err != nil {
		fmt.Println("Error reading CSV file:", err)
		return
	}

	correct := 0
	for i, q := range questions {
		select {
		case <-timeoutCh:
			fmt.Println("Time limit reached!")
			fmt.Printf("You answered %d out of %d questions correctly.\n", correct, len(questions))
			return
		default:
			if askQuestion(i+1, q) {
				fmt.Println("Correct!")
				correct++
			} else {
				fmt.Printf("Incorrect! The correct answer: %s\n", q.answer)
			}
		}
	}

	fmt.Printf("You answered %d out of %d questions correctly.\n", correct, len(questions))
}
