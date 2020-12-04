package main

import (
	"../utils"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	lines := utils.ReadLines("./day04/puzzle_input.txt")

	validPassps := 0
	allFieldsPresentCnt := 0
	var data = make(map[string]string)
	for _, line := range lines {
		if line != "" && line != "\n" {
			splitLine := strings.Split(line, " ")
			for _, spl := range splitLine {
				kv := strings.Split(spl, ":")
				data[kv[0]] = kv[1]
			}
		} else {
			val, allFieldsPresent := validate(data)
			if allFieldsPresent {
				allFieldsPresentCnt++
			}
			if val {
				validPassps++
				fmt.Println(data)
			}
			data = make(map[string]string)
		}
	}

	val, allFieldsPresent := validate(data)
	if allFieldsPresent {
		allFieldsPresentCnt++
	}
	if val {
		validPassps++
		fmt.Println(data)
	}

	log.Println("Result Part1: ", allFieldsPresentCnt)
	log.Println("Result Part2: ", validPassps)
}

func validate(data map[string]string) (bool, bool) {
	valid := true
	allFieldsPresent := true
	for _, c := range []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"} {
		_, ok := data[c]
		if !ok {
			allFieldsPresent = false
			valid = false
			break
		}
	}
	if allFieldsPresent {
		byr, _ := strconv.Atoi(data["byr"])
		if !(byr >= 1920 && byr <= 2002) {
			valid = false
		}
		iyr, _ := strconv.Atoi(data["iyr"])
		if !(iyr >= 2010 && iyr <= 2020) {
			valid = false
		}
		eyr, _ := strconv.Atoi(data["eyr"])
		if !(eyr >= 2020 && eyr <= 2030) {
			valid = false
		}
		hgt := data["hgt"]
		if strings.HasSuffix(hgt, "cm") {
			hsplit := strings.Split(hgt, "c")
			hgtval, _ := strconv.Atoi(hsplit[0])
			if !(hgtval >= 150 && hgtval <= 193) {
				valid = false
			}
		} else if strings.HasSuffix(hgt, "in") {
			hsplit := strings.Split(hgt, "i")
			hgtval, _ := strconv.Atoi(hsplit[0])
			if !(hgtval >= 59 && hgtval <= 76) {
				valid = false
			}
		} else {
			valid = false
		}
		hcl := data["hcl"]
		hclmatch, _ := regexp.Match("^#[a-f0-9]{6}$", []byte(hcl))
		if !hclmatch {
			valid = false
		}
		ecl := data["ecl"]
		eyclValid, _ := regexp.MatchString("^amb|blu|brn|gry|grn|hzl|oth$", ecl)
		if !eyclValid {
			valid = false
		}
		pid := data["pid"]
		pidmatch, _ := regexp.Match("^[0-9]{9}$", []byte(pid))
		if !pidmatch {
			valid = false
		}
	}
	return valid, allFieldsPresent
}
