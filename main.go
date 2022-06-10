package main

import "flag"

func main() {
	//create name for csv file flagname, default name, help text
	csvFilename := flag.String("csv", "problems.csv", "A csv file in the format 'question,answer'")
	flag.Parse()

	// go build main.go
	// go run main.go at this point gives no results
	// however: go run main.go -h displays the csvFilename parsed in the terminal
	_ = csvFilename
	// ^ only doing this so the code compiles, unused variables will stop code compilation
}
