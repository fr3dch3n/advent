package day01

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

var DIGIT_LOOKUP = map[string]int{
	"zero":  0,
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

type FoundDigit struct {
	idx, digit int
	spelledOut bool
}

func extractAllDigits(line string) []FoundDigit {
	var result []FoundDigit
	lineLen := len(line)
	for i := 0; i < lineLen; i++ {
		if unicode.IsDigit(rune(line[i])) {
			result = append(result, FoundDigit{i, int(line[i] - '0'), false})
		}
		for j := i + 1; j <= lineLen; j++ {
			start, digit := subStringIsSpelledOutDigit(line, i, j)
			if start != -1 {
				result = append(result, FoundDigit{start, digit, true})
			}
		}
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].idx < result[j].idx
	})
	return result
}

func subStringIsSpelledOutDigit(line string, start, end int) (int, int) {
	s := line[start:end]
	spelledOutDigits := make([]string, 0, len(DIGIT_LOOKUP))
	for k := range DIGIT_LOOKUP {
		spelledOutDigits = append(spelledOutDigits, k)
	}
	for _, spelledOutDigit := range spelledOutDigits {

		res1 := strings.Index(s, spelledOutDigit)
		if res1 != -1 {
			return res1 + start, DIGIT_LOOKUP[spelledOutDigit]
		}
	}
	return -1, -1
}

func extractFirstAndLastDigit(line string, spelledOut bool) int {
	var firstDigit, lastDigit string
	foundDigits := extractAllDigits(line)

	if !spelledOut {
		nonSpelledOutDigits := make([]FoundDigit, 0, len(foundDigits))
		for _, foundDigit := range foundDigits {
			if !foundDigit.spelledOut {
				nonSpelledOutDigits = append(nonSpelledOutDigits, FoundDigit{foundDigit.idx, foundDigit.digit, false})
			}
		}
		firstDigit = fmt.Sprint(nonSpelledOutDigits[0].digit)
		lastDigit = fmt.Sprint(nonSpelledOutDigits[len(nonSpelledOutDigits)-1].digit)
	} else {
		firstDigit = fmt.Sprint(foundDigits[0].digit)
		lastDigit = fmt.Sprint(foundDigits[len(foundDigits)-1].digit)
	}
	combined := firstDigit + lastDigit
	i, _ := strconv.Atoi(combined)
	return i
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
		return extractFirstAndLastDigit(line, false)
	})
}
func solutionB() int {
	return runFunctionForEachLineAndSumResult(func(line string) int {
		return extractFirstAndLastDigit(line, true)
	})
}
