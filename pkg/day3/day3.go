package day3

import (
	"bufio"
	"log"
	"os"
)

type Object int

const (
	Tree Object = iota
	Empty
)

type Map [][]Object

func Day3() error {
	m, err := readInput("assets/inputs/day3.txt")
	if err != nil {
		return err
	}
	count := countTrees(m, 3, 1)
	log.Printf("day3 part1 output: %v", count)
	part2 := countTrees(m, 1, 1) * countTrees(m, 3, 1) * countTrees(m, 5, 1) * countTrees(m, 7, 1) * countTrees(m, 1, 2)
	log.Printf("day3 part2 output: %v", part2)
	return nil
}

func countTrees(m Map, xslope int, yslope int) int {
	xmax := len(m[0])
	ymax := len(m)
	x, y := 0, 0
	count := 0
	for y < ymax {
		if m[y][x] == Tree {
			count++
		}
		y += yslope
		x = (x + xslope) % xmax
	}
	return count
}

func readInput(path string) (Map, error) {
	result := make(Map, 0)
	inputFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(inputFile)
	y := 0
	var obj Object
	for scanner.Scan() {
		result = append(result, make([]Object, 0))
		for _, token := range scanner.Bytes() {
			switch token {
			case '.':
				obj = Empty
			case '#':
				obj = Tree
			default:
				log.Fatalf("unexpected token at y=%v: '%v'", y, token)
			}
			result[y] = append(result[y], obj)
		}
		y++
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return result, nil
}
