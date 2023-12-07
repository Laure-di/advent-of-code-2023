package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldReturnAMapWith2Pair02And49_Whentwo1nine(t *testing.T) {
	expected := map[int]int{
		0: 2,
		4: 9,
	}
	expectedIndex := 0
	index := findSmallestKey(expected)
	answer := findFirstOccurrenceLetterDigit("two1nine")
	assert.Equal(t, expected, answer)
	assert.Equal(t, expectedIndex, index)
}

func TestShouldReturnAMapWith3Pair08And42And73_Wheneightwothree(t *testing.T) {
	expected := map[int]int{
		0: 8,
		4: 2,
		7: 3,
	}
	answer := findFirstOccurrenceLetterDigit("eightwothree")
	assert.Equal(t, expected, answer)
	expectedIndex := 7
	index := findLargestKey(expected)
	assert.Equal(t, expectedIndex, index)
}

func TestShouldReturnAMapWith3Pair19And58And107_When4nineeightseven2(t *testing.T) {
	expected := map[int]int{
		1:  9,
		5:  8,
		10: 7,
	}
	answer := findFirstOccurrenceLetterDigit("4nineeightseven2")
	assert.Equal(t, expected, answer)
	expectedIndex := 1
	index := findSmallestKey(expected)
	assert.Equal(t, expectedIndex, index)
}

func TestShouldReturnAMapWith2Pair31And73_Whenabcone2threexyz(t *testing.T) {
	expected := map[int]int{
		3: 1,
		7: 3,
	}
	answer := findFirstOccurrenceLetterDigit("abcone2threexyz")
	assert.Equal(t, expected, answer)
	expectedIndex := 3
	index := findSmallestKey(expected)
	assert.Equal(t, expectedIndex, index)
}

func TestShouldReturnAMapWith3Pair12And31And74_Whenxtwone3four(t *testing.T) {
	expected := map[int]int{
		1: 2,
		3: 1,
		7: 4,
	}
	answer := findFirstOccurrenceLetterDigit("xtwone3four")
	assert.Equal(t, expected, answer)
	expectedIndex := 1
	index := findSmallestKey(expected)
	assert.Equal(t, expectedIndex, index)
}
