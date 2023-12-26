package day04

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Card struct {
	Id             int
	WinningNumbers []int
	MyNumbers      []int
	Copies         int
}

func (c *Card) Points() int {
	matches := 0
	for _, v := range c.MyNumbers {
		for _, w := range c.WinningNumbers {
			if v == w {
				matches++
			}
		}
	}
	return matches
}

func numbersString(numbers string) []int {
	n_parts := strings.Split(numbers, " ")
	var n []int
	for _, v := range n_parts {
		if v == "" {
			continue
		}
		vv := strings.Trim(v, " ")
		v_int, err := strconv.Atoi(vv)
		if err != nil {
			panic(err)
		}
		n = append(n, v_int)
	}
	return n
}

func lineToCard(line string) Card {
	re := regexp.MustCompile(`Card\s+(\d+): (.+) \| (.+)`)
	results := re.FindStringSubmatch(line)
	if len(results) != 4 {
		fmt.Printf("Invalid line: %s\n", line)
		panic("Invalid line")
	}
	idRawTrimmed := strings.Trim(results[1], " ")
	id, _ := strconv.Atoi(idRawTrimmed)
	winningNumbers := numbersString(results[2])
	myNumbers := numbersString(results[3])

	return Card{Id: id, WinningNumbers: winningNumbers, MyNumbers: myNumbers, Copies: 1}
}

func solutionA() int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		card := lineToCard(scanner.Text())
		matches := card.Points()
		sum += int(math.Pow(2, float64(matches-1)))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return sum
}

func processCardCopying(cards []Card) int {
	for i, card := range cards {
		points := card.Points()
		for x := 0; x < card.Copies; x++ {
			for j := i + 1; j <= i+points; j++ {
				if j >= len(cards) {
					break
				}
				cards[j].Copies++
			}
		}
	}
	sumCopies := 0
	for _, card := range cards {
		sumCopies += card.Copies
	}
	return sumCopies
}

func solutionB() int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var cards []Card
	for scanner.Scan() {
		card := lineToCard(scanner.Text())
		cards = append(cards, card)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return processCardCopying(cards)
}
