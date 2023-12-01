package main

import (
	"aoc2023/day1"
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)
func main() {
	dayFlag := flag.Int("d", 0, "Day to run")

	flag.Parse()

	reader := bufio.NewReader(os.Stdin)
	inputLines := make([]string, 0)
	for {
		line, err := reader.ReadString('\n')
		inputLines = append(inputLines, line)
		if err != nil {
			break
		}
	}

	writer := bufio.NewWriter(os.Stdout)

	switch day := *dayFlag; day {
	case 1:
		writer.WriteString(fmt.Sprintln(day1.Sum(inputLines)))
		writer.Flush()
	default:
		log.Fatal("Invalid day")
	}
}