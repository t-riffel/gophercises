package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func main() {
	// quiz problems passed in through command-line
	args := os.Args[1:]

	// Open the file and parse quiz questions
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

		s := strings.Split(record[0], ",")
		q, a := s[0], s[1]

		// Store the question/answer
		questions[count] = q
		answers[count] = a
		count++
	}

	t1 := time.NewTimer(30 * time.Second)
	done := make(chan bool)

	go func() {
		for i := 0; i < count; i++ {
			fmt.Println("What is", questions[i], "?")

			// Wait for answer
			r := bufio.NewReader(os.Stdin)
			text, _ := r.ReadString('\n')

			// Check if input matches answer
			if strings.TrimRight(text, "\n") == answers[i] {
				score++
			}
		}
		done <- true
	}()

	select {
	case <-done:
	case <-t1.C:
		fmt.Println("Time's up!")
	}

	fmt.Println("Final Score:", score, "/", count)
}
