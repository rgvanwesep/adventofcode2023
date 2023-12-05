package main

import (
	"aoc2023/day1"
	"aoc2023/day2"
	"aoc2023/day3"
	"aoc2023/day4"
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	dayFlag := flag.Int("d", 0, "Day to run")
	partFlag := flag.Int("p", 1, "Part to run")

	flag.Parse()

	day, part := *dayFlag, *partFlag
	log.Printf("Running day %d, part %d", day, part)

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

	switch [2]int{day, part} {
	// This case isn't quite right, because only day 1 part 2 is implemented
	case [2]int{1, 1}:
		writer.WriteString(fmt.Sprintln(day1.Sum(inputLines)))
	case [2]int{2, 1}:
		writer.WriteString((fmt.Sprintln(day2.Sum(inputLines))))
	case [2]int{2, 2}:
		writer.WriteString((fmt.Sprintln(day2.SumPower(inputLines))))
	case [2]int{3, 1}:
		writer.WriteString((fmt.Sprintln(day3.Sum(inputLines))))
	case [2]int{3, 2}:
		writer.WriteString((fmt.Sprintln(day3.SumGearRatios(inputLines))))
	case [2]int{4, 1}:
		writer.WriteString((fmt.Sprintln(day4.Sum(inputLines))))
	case [2]int{4, 2}:
		writer.WriteString((fmt.Sprintln(day4.SumCards(inputLines))))
	default:
		log.Fatal("Invalid day or part")
	}

	writer.Flush()
}
