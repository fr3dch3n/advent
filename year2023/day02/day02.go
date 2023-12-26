package day02

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Set struct {
	Blues  int
	Reds   int
	Greens int
}
type Game struct {
	Id   int
	Sets []Set
}
type BagConfiguration struct {
	Blues  int
	Reds   int
	Greens int
}

func parseLineToGame(line string) Game {
	parts := strings.Split(line, ": ")
	id_part := parts[0]
	id_str, _ := strings.CutPrefix(id_part, "Game ")
	id, _ := strconv.Atoi(id_str)

	sets_part := parts[1]
	var sets []Set
	for _, set := range strings.Split(sets_part, ";") {
		color_parts := strings.Split(set, ",")
		var blues, reds, greens int
		for _, color_part := range color_parts {
			cp := strings.TrimSpace(color_part)
			if strings.Contains(cp, "blue") {
				amount_blues, _ := strings.CutSuffix(cp, " blue")
				blues, _ = strconv.Atoi(amount_blues)
			}
			if strings.Contains(cp, "red") {
				amount_reds, _ := strings.CutSuffix(cp, " red")
				reds, _ = strconv.Atoi(amount_reds)
			}
			if strings.Contains(cp, "green") {
				amount_greens, _ := strings.CutSuffix(cp, " green")
				greens, _ = strconv.Atoi(amount_greens)
			}
		}
		sets = append(sets, Set{blues, reds, greens})
	}
	return Game{Id: id, Sets: sets}
}

func gameIsValidWithBagConfiguration(game Game, bagConfiguration BagConfiguration) bool {
	for _, set := range game.Sets {
		if set.Blues > bagConfiguration.Blues || set.Reds > bagConfiguration.Reds || set.Greens > bagConfiguration.Greens {
			return false
		}
	}
	return true
}

func minNecessaryBagConfigurationForGame(game Game) BagConfiguration {
	var blues, reds, greens int
	for _, set := range game.Sets {
		if set.Blues > blues {
			blues = set.Blues
		}
		if set.Reds > reds {
			reds = set.Reds
		}
		if set.Greens > greens {
			greens = set.Greens
		}
	}
	return BagConfiguration{Blues: blues, Reds: reds, Greens: greens}
}

func runFunctionForEachLineAndSumResult(f func(string) int) int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sum += f(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return sum
}

func solutionA() int {
	return runFunctionForEachLineAndSumResult(func(line string) int {
		bagConfiguration := BagConfiguration{Blues: 14, Reds: 12, Greens: 13}
		game := parseLineToGame(line)
		if isValid := gameIsValidWithBagConfiguration(game, bagConfiguration); isValid {
			return game.Id
		}
		return 0
	})
}

func solutionB() int {
	return runFunctionForEachLineAndSumResult(func(line string) int {
		game := parseLineToGame(line)
		minBagConfiguration := minNecessaryBagConfigurationForGame(game)
		power := minBagConfiguration.Blues * minBagConfiguration.Reds * minBagConfiguration.Greens
		return power
	})
}
