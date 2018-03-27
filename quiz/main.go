package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func main() {
	var csvPath = flag.String("csv", "problems.csv", "the csv file from gophercises")
	var timeLimit = flag.Int("time", 30, "time limit for the whole quiz in seconds")

	flag.Parse()

	// Open the file
	file, err := os.Open(*csvPath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	c := csv.NewReader(file)
	c.Comma = ';'

	score := 0
	count := 0
	// Max number of questions will be 100
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

	fmt.Println("Press ENTER to start the quiz/timer")
	bufio.NewScanner(os.Stdout).Scan()

	t1 := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	done := make(chan bool)

	go func() {
		for i := 0; i < count; i++ {
			fmt.Println("What is", questions[i], "?")

			// Wait for answer
			r := bufio.NewReader(os.Stdin)
			text, _ := r.ReadString('\n')

			// Check if input matches answer
			if strings.TrimRight(strings.TrimSpace(strings.ToLower(text)), "\n") == answers[i] {
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
