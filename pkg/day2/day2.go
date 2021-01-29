package day2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Day2() error {
	pwlist, err := readInput("./assets/inputs/day2.txt")
	if err != nil {
		return err
	}
	oldPolicyMatches := 0
	newPolicyMatches := 0
	for _, pw := range pwlist {
		if pw.MatchesOldPolicy() {
			oldPolicyMatches++
		}
		if pw.MatchesNewPolicy() {
			newPolicyMatches++
		}
	}
	log.Printf("day2 part1 output: %v", oldPolicyMatches)
	log.Printf("day2 part2 output: %v", newPolicyMatches)
	return nil
}

type Range struct {
	min int
	max int
}

type PasswordInfo struct {
	MandatoryChar  byte
	OccurenceRange Range
	Password       string
}

func (pwinfo *PasswordInfo) MatchesOldPolicy() bool {
	count := strings.Count(pwinfo.Password, string(pwinfo.MandatoryChar))
	return (count >= pwinfo.OccurenceRange.min && count <= pwinfo.OccurenceRange.max)
}

func (pwinfo *PasswordInfo) MatchesNewPolicy() bool {
	firstChar := pwinfo.Password[pwinfo.OccurenceRange.min-1]
	lastChar := pwinfo.Password[pwinfo.OccurenceRange.max-1]
	if (firstChar == pwinfo.MandatoryChar && lastChar != pwinfo.MandatoryChar) ||
		(firstChar != pwinfo.MandatoryChar && lastChar == pwinfo.MandatoryChar) {
		return true
	}
	return false
}

func ParsePasswordInfo(inputLine string) (*PasswordInfo, error) {
	r := regexp.MustCompile(`(?P<min>\d+)-(?P<max>\d+) (?P<char>[a-z]): (?P<password>\w+)`)
	matches := r.FindStringSubmatch(inputLine)
	if matches == nil {
		return nil, fmt.Errorf("can not parse input line '%s'", inputLine)
	}
	min, err := strconv.Atoi(matches[1])
	if err != nil {
		log.Fatal(err)
	}
	max, err := strconv.Atoi(matches[2])
	if err != nil {
		log.Fatal(err)
	}
	char := matches[3][0]
	password := matches[4]
	occurences := Range{min, max}
	return &PasswordInfo{
		char,
		occurences,
		password,
	}, nil
}

func readInput(path string) ([]PasswordInfo, error) {
	result := make([]PasswordInfo, 0)
	inputFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		pwinfo, err := ParsePasswordInfo(scanner.Text())
		if err != nil {
			return nil, err
		}
		result = append(result, *pwinfo)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return result, nil
}
