package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"unicode"
)

func findFirstOccurrenceLetterDigit(text string) map[int]int {
	letter := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	i := 0
	answer := map[int]int{}
	for _, digit := range letter {

		re := regexp.MustCompile(digit)

		index := re.FindStringIndex(text)
		i++

		if index != nil {
			answer[index[0]] = i
		}
	}
	return answer
}

func findLastOccurrenceLetterDigit(text string) map[int]int {
	letter := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	i := 0
	answer := map[int]int{}
	for _, digit := range letter {

		re := regexp.MustCompile(digit)

		index := re.FindAllStringIndex(text, -1)
		i++

		if len(index) > 0 {
			lastOccurencePosition := index[len(index)-1]
			answer[lastOccurencePosition[0]] = i
		}
	}
	return answer
}

func findSmallestKey(myMap map[int]int) int {
	smallestKey := -1
	for key := range myMap {
		if key < smallestKey || smallestKey == -1 {
			smallestKey = key
		}
	}

	return smallestKey
}

func findLargestKey(myMap map[int]int) int {
	largestKey := 0
	for key := range myMap {
		if largestKey == 0 || key > largestKey {
			largestKey = key
		}
	}

	return largestKey
}

func calculateCalibration(text string) int {
	number := 0
	answerFirstOccurence := findFirstOccurrenceLetterDigit(text)
	answerLastOccurence := findLastOccurrenceLetterDigit(text)
	smallestKey := findSmallestKey(answerFirstOccurence)
	largestKey := findLargestKey(answerLastOccurence)

	number = findFirstDigit(text, smallestKey, answerFirstOccurence)
	number += findLastDigit(text, largestKey, answerLastOccurence)
	println(number)
	return number
}

func findFirstDigit(text string, smallestKey int, answerFirstOccurence map[int]int) int {
	for y := 0; y < len(text); y++ {
		if unicode.IsDigit(rune(text[y])) || smallestKey < y && smallestKey != -1 {
			letterDigit, ok := answerFirstOccurence[smallestKey]
			if y < smallestKey || !ok {
				return int(text[y]-'0') * 10
			} else {
				return letterDigit * 10
			}
			break
		}
	}
	return 0
}

func findLastDigit(text string, largestKey int, answerLastOccurence map[int]int) int {
	for i := len(text) - 1; i >= 0; i-- {
		if unicode.IsDigit(rune(text[i])) || i < largestKey && largestKey != -1 {
			letterDigit, ok := answerLastOccurence[largestKey]
			if largestKey < i || !ok {
				return int(text[i] - '0')
			} else {
				return letterDigit
			}
			break
		}
	}
	return 0
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	number := 0
	for fileScanner.Scan() {
		number += calculateCalibration(fileScanner.Text())
	}
	file.Close()
}
