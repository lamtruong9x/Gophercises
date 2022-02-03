package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {
	csvFileName := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	flag.Parse() // turning string into a data structure in program
	_ = csvFileName

	file, err := os.Open(*csvFileName)
	defer file.Close()
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFileName))
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		fmt.Printf("Can read the file")
	}
	fmt.Println(lines)
	questionbank := []problem{}
	for i, e := range lines {
		questionbank[i] = problem{
			question: e[0],
			answer: e[1],
		}
	}
	fmt.Println(questionbank)
}

type problem struct {
	question string
	answer string
}

func exit(mess string) {
	fmt.Println(mess)
	os.Exit(1)
}