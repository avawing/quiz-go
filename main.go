package main

import (
	"flag"
	"fmt"
	"os"
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

	_ = file
	// no unused variables for code compilation
}

// functions have one purpose

func exit(msg string) {
	fmt.Println(msg)
	// something went wrong - status 500 equivalent
	os.Exit(1)
}
