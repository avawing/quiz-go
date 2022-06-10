package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func main() {
	//create name for csv file flag name, default name, help text
	// csvFilename is a pointer to a string, this is just how the flag package works
	csvFilename := flag.String("csv", "problems.csv", "A csv file in the format 'question,answer'")
	timeLimit := flag.Int("limit", 30, "The time limit for the quiz in seconds")
	flag.Parse()

	// go build main.go
	// go run main.go at this point gives no results
	// however: go run main.go -h displays the csvFilename parsed in the terminal

	// because flag package is a pointer to a flag name, we have to use it as a pointer in os.Open
	// os.Open returns a file and an error
	file, err := os.Open(*csvFilename)

	//os.Open will only return an error if the file does not exist / a problem occurs
	// check that error is not nil
	if err != nil {
		// csvFilename is a pointer to a string -> %s is string, *csvFilename -> pointer
		// passing format as a variable to the exit function
		exit(fmt.Sprintf("failed to open the csv file: %s ", *csvFilename))
	}

	// reader / writers are very common in go
	// setup work THEN start timer
	lines := createLines(file)
	problems := parseLines(lines)
	// create timer -> fixed *timeLimit is not time.Duration by coercing type
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	// initialize counter
	correct := 0
	// index, problem
	for i, p := range problems {
		fmt.Printf("Problem %d: %s = \n", i+1, p.q)
		answerChan := make(chan string)
		// anonymous function
		// go routine
		go func() {
			var answer string
			_, err := fmt.Scanf("%s", &answer)
			if err != nil {
				exit("Oops!")
			}
			answerChan <- answer
		}()
		select {
		case <-timer.C:
			fmt.Printf("You scored %d out of %d \n", correct, len(problems))
			//breaks fully out of for loop
			return
		case answer := <-answerChan:
			if answer == p.a {
				correct++
			}
		}
		// account for zero indexing, access question

	}
	fmt.Printf("You scored %d out of %d \n", correct, len(problems))
}

// type

type problem struct {
	q string
	a string
}

func createLines(file io.Reader) [][]string {
	r := csv.NewReader(file)

	// parse csv -> entire file upfront -> small file, not going to cause mem issues
	lines, err := r.ReadAll()

	if err != nil {
		exit("failed to parse provided csv file")
	}
	return lines
}
func parseLines(lines [][]string) []problem {
	// make array of problems that is the length of the lines
	// arrays have predefined length!!
	// faster to create 'bucket' in this way
	ret := make([]problem, len(lines))
	// for index, line in range of lines
	for i, line := range lines {
		// create a problem (LOL)
		// trailing comma necessary
		ret[i] = problem{
			q: line[0],
			// trim spaces from answers - in case of poor formatting
			a: strings.TrimSpace(line[1]),
		}
	}
	return ret
}

// functions have one purpose

func exit(msg string) {
	fmt.Println(msg)
	// something went wrong - status 500 equivalent
	os.Exit(1)
}
