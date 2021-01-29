package day1

import (
	"bufio"
	"errors"
	"log"
	"os"
	"sort"
	"strconv"
)

func Day1() error {
	input, err := readInput("assets/inputs/day1.txt")
	if err != nil {
		return err
	}
	sort.Ints(input)
	i, j, err := findTwoValuesWithSum(input, 2020)
	if err != nil {
		return err
	}
	log.Printf("day1 part1 output: %d (from %d %d)", i*j, i, j)
	i, j, k, err := findThreeValuesWithSum(input, 2020)
	if err != nil {
		return err
	}
	log.Printf("day1 part2 output: %d (from %d %d %d)", i*j*k, i, j, k)
	return nil
}

func findThreeValuesWithSum(list []int, sum int) (int, int, int, error) {
	for _, lower := range list {
		for j, _ := range list {
			upper1 := list[len(list)-j-1]
			for k, _ := range list {
				upper2 := list[len(list)-k-1]
				if upper1+upper2+lower == sum {
					return upper1, upper2, lower, nil
				}
				if upper1+upper2+lower < sum {
					break
				}
			}
		}
	}
	return 0, 0, 0, errors.New("could not find two values with the given sum")
}

func findTwoValuesWithSum(list []int, sum int) (int, int, error) {
	for _, lower := range list {
		for j, _ := range list {
			upper := list[len(list)-j-1]
			if upper+lower == sum {
				return upper, lower, nil
			}
			if upper+lower < sum {
				break
			}
		}
	}
	return 0, 0, errors.New("could not find two values with the given sum")
}

func readInput(path string) ([]int, error) {
	result := make([]int, 0)
	inputFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		result = append(result, i)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return result, nil
}
