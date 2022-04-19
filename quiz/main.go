package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	// Flags that we can pass to the app like ./quiz.exe -csv=file123.csv -limit=2
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	timeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds")
	flag.Parse()
	
	// Opens the file
	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFilename))
	}
	
	r := csv.NewReader(file)
	// Read all the lines. No problem because our file is small
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the CSV file")
	}

	problems := parseLines(lines)

	// Create a time that sends a 'signal' to channel C after its Duration https://pkg.go.dev/time#NewTimer
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	correct := 0
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, p.Q)
		
		// Channel that will receive the anser from the user
		// We used this Goroutine and the channel because this way we can end the quiz when the timer is over
		// even when the user is 'idle'. We don't get stuck in the Scanf event
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()

		// "Listens" the channels to know what to do
		select {
		// Timer is over. Channel C receives the 'signal'
		case <- timer.C:
			fmt.Printf("\nYou scored %d out of %d\n", correct, len(problems))
			return
		// User typed an answer and we get it from answerCh
		case answer := <- answerCh:
			if answer == p.A {
				correct++
			}
		}
	}
	
	fmt.Printf("\nYou scored %d out of %d\n", correct, len(problems))
}

type Problem struct {
	Q string
	A string
}

func parseLines(lines [][]string) []Problem {
	// Create a slice of Problem with the right size
	// We know that each line of the file is a question, so we know how many questions there are
	ret := make([]Problem, len(lines))

	// range on slices returns the index and the value for each entry
	for i, line := range lines {
		ret[i] = Problem {
			Q: line[0],
			A: strings.TrimSpace(line[1]),
		}
	}

	return ret
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}