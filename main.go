package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	//create name for csv file flag name, default name, help text
	// csvFilename is a pointer to a string, this is just how the flag package works
	csvFilename := flag.String("csv", "problems.csv", "A csv file in the format 'question,answer'")
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
	r := csv.NewReader(file)

	// parse csv -> entire file upfront -> small file, not going to cause mem issues
	lines, err := r.ReadAll()

	if err != nil {
		exit("failed to parse provided csv file")
	}

	problems := parseLines(lines)

	// initialize counter
	correct := 0
	// index, problem
	for i, p := range problems {
		// account for zero indexing, access question
		fmt.Printf("Problem %d: %s = \n", i+1, p.q)
		var answer string
		// scan for entered string (Scan f removed spaces), &answer points to variable and updates it
		_, err := fmt.Scanf("%s\n", &answer)
		if err != nil {
			exit("Something hella broke")
		}
		// check correctness of answer
		if answer == p.a {
			correct++
		}
	}
	fmt.Printf("You scored %d out of %d \n", correct, len(problems))
}

// type

type problem struct {
	q string
	a string
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
