package main

import (
	"aoc2023/day1"
	"aoc2023/day2"
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
	log.Printf("Read %d lines\n", len(inputLines))

	writer := bufio.NewWriter(os.Stdout)

	switch day := *dayFlag; day {
	case 1:
		writer.WriteString(fmt.Sprintln(day1.Sum(inputLines)))
	case 2:
		writer.WriteString((fmt.Sprintln(day2.Sum(inputLines))))
	default:
		log.Fatal("Invalid day")
	}

	writer.Flush()
}
