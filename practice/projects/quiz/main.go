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
	csvFilename := flag.String("csv", "problems.csv", "A csv fle in the format of 'number to add, answer'")
	timeLimit := flag.Int("limit", 30, "The time limit for the quiz in seconds")
	flag.Parse()

	file, err := os.Open(*csvFilename)

	if err != nil {
		exit(fmt.Sprintf("Cannot open file: %s", *csvFilename))
		os.Exit(1)
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to read the CSV file")
	}
    problems := parseLines(lines)
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	correct := 0
	problemloop:
	for i, p := range problems{
		answerCh := make(chan string)
		go func() {			
			fmt.Printf("Problem #%d: %s = \n", i+1 ,  p.q)	
			var answer string;
			fmt.Scanf("%s\n", &answer)  
			answerCh <- answer
		}()
		select {
		case <-timer.C:
			fmt.Printf("You got %d questions correct.\n", correct)
			break problemloop
		case answer := <-answerCh:
			if answer == p.a {
				correct++
			}
		}

	}
	fmt.Printf("You got %d questions correct.\n", correct)

}

func parseLines(lines [][]string) []problem {
 ret := make([]problem, len(lines))
 fmt.Println(len(lines))
 for i, line := range lines{
	ret[i] = problem{
		q: line[0],
		a: strings.TrimSpace(line[1]),
	}
 }
 return ret
}

type problem struct {
	q string
	a string
}

func exit(msg string){
	fmt.Println(msg)
	os.Exit(1)
} 