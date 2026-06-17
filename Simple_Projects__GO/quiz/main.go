package main 

import (
	"fmt"
	"os"
	"encoding/csv"
	"flag"
	"strings"
)

// problem struct
type problem struct {
	q string 
	a string 
}

func main() {
	// define flags 
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question, answer' ")
	flag.Parse()

	// Open the csv file 
	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the cse file: %s", *csvFilename))

	}
	// close the file after execution
	defer file.Close()
	
	// Read and parse the csv content 
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("failed to pase the provided csv file")

	}
	problems := parseLines(lines)
	// quiz loop 
	correct := 0 

	for i, p := range problems {
		fmt.Printf("Problems #%d: %s = = ", i+1, p.q)

		var answer string 
		fmt.Scanf("%s\n", &answer)

		if answer == p.a {
			correct++ 
		}
	}
	fmt.Printf("\nYou scored %d out of %d,\n", correct, len(problems))
}


// parseLines converts a 2d slice of strings into a slice of problem structs 
func parseLines(lines [][]string) []problem {

	ret := make([]problem, len(lines))


	for i, line := range lines {
		ret[i] = problem {
			q: line[0], 
			a: strings.TrimSpace(line[1]),
		}
	}
	return ret 
}

// exit prints an error message and terminates the program 
func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
