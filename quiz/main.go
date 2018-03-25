package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	// quiz problems passed in through command-line
	args := os.Args[1:]

	// Open the file and parse quiz questions, maybe use a map
	file, err := os.Open(args[0])
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	c := csv.NewReader(file)

	c.Comma = ';'
	score := 0
	count := 0
	var questions, answers [100]string

	for {
		// Read a single record which is one line
		record, err := c.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error:", err)
			return
		}

		// Split the record
		s := strings.Split(record[0], ",")
		q, a := s[0], s[1]

		// Store the question/answer
		questions[count] = q
		answers[count] = a
		count++
	}

	for i := 0; i < count; i++ {
		// Print question
		fmt.Println("What is", questions[i], "?")

		// Wait for answer
		r := bufio.NewReader(os.Stdin)
		text, _ := r.ReadString('\n')

		// Check if input matches answer
		if strings.TrimRight(text, "\n") == answers[i] {
			score++
		}
	}

	// Print final score
	fmt.Println("Final Score:", score, "/", count)
}
