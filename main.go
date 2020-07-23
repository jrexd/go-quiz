package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

type problem struct {
	q string
	a string
}

func main() {
	var path string
	flag.StringVar(&path, "p", "", "Path to .csv file")
	flag.Parse()

	var wrong int = 0
	var correct int = 0

	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file.")
	}

	data := csv.NewReader(file)
	res, err := data.ReadAll()
	if err != nil {
		fmt.Println("Error parsing csv.")
	}

	problems := parseLines(res)

	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.q)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer != p.a {
			wrong++
		} else {
			correct++
		}
	}
	fmt.Printf("Correct: %d\nWrong: %d\n", correct, wrong)
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: line[1],
		}
	}
	return ret
}
