package day4

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Day4() error {
	passports, err := readInput("assets/inputs/day4.txt")
	if err != nil {
		return err
	}
	hasRequiredFieldsCount := 0
	isValidCount := 0
	for _, password := range passports {
		if password.hasRequiredFields() {
			hasRequiredFieldsCount++
		}
		if password.isValid() {
			isValidCount++
		}
	}
	log.Printf("day4 part1 output: %v", hasRequiredFieldsCount)
	log.Printf("day4 part2 output: %v", isValidCount)
	return nil
}

func (passport *Passport) hasRequiredFields() bool {
	requiredFields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	for _, field := range requiredFields {
		_, ok := (*passport)[field]
		if !ok {
			return false
		}
	}
	return true
}

func validateDate(field string, min int, max int) bool {
	i, err := strconv.Atoi(field)
	if err != nil {
		return false
	}
	return i >= min && i <= max
}

func validateHeight(field string) bool {
	r := regexp.MustCompile(`^(?P<height>\d{2,3})(?P<unit>(cm|in))$`)
	matches := r.FindStringSubmatch(field)
	if matches == nil {
		return false
	}
	height, err := strconv.Atoi(matches[1])
	if err != nil {
		return false
	}
	unit := matches[2]
	switch unit {
	case "cm":
		return height >= 150 && height <= 193
	case "in":
		return height >= 59 && height <= 76
	default:
		return false
	}
}

func validateHairColor(field string) bool {
	r := regexp.MustCompile(`^#[0-9a-z]{6}$`)
	return r.MatchString(field)
}

func validateEyeColor(field string) bool {
	return field == "amb" || field == "blu" || field == "brn" || field == "gry" || field == "grn" || field == "hzl" || field == "oth"
}

func validatePassportId(field string) bool {
	r := regexp.MustCompile(`^[0-9]{9}$`)
	return r.MatchString(field)
}

func (passport *Passport) isValid() bool {
	if !passport.hasRequiredFields() {
		return false
	}
	return validateDate((*passport)["byr"], 1920, 2002) &&
		validateDate((*passport)["iyr"], 2010, 2020) &&
		validateDate((*passport)["eyr"], 2020, 2030) &&
		validateHeight((*passport)["hgt"]) &&
		validateHairColor((*passport)["hcl"]) &&
		validateEyeColor((*passport)["ecl"]) &&
		validatePassportId((*passport)["pid"])

}

type Passport map[string]string

func readInput(path string) ([]Passport, error) {
	result := make([]Passport, 0)
	inputFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(inputFile)
	passport := make(Passport)
	for scanner.Scan() {
		line := scanner.Text()
		// complete passport, start reading next one
		if len(line) == 0 {
			result = append(result, passport)
			passport = make(Passport)
			continue
		}
		fields := strings.Fields(line)
		for _, field := range fields {
			items := strings.Split(field, ":")
			if len(items) != 2 {
				log.Fatalf("got '%v', expected something with the format 'key:value', while parsing line '%v'", field, line)
			}
			passport[items[0]] = items[1]
		}
	}
	result = append(result, passport)
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return result, nil
}
