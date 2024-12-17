package internal

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func LoadFile(path string) string {
	fileByte, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}

	fileString := string(fileByte)

	return fileString
}

type Mul struct {
	firstNumber  int
	secondNumber int
}

func ExtractMulNumbers(stringContent string) []Mul {
	validMuls := []Mul{}

	re := regexp.MustCompile(`mul\((?P<firstNumber>[0-9]{1,3}),(?P<secondNumber>[0-9]{1,3})\)`)
	n1 := re.SubexpNames()
	matches := re.FindAllStringSubmatch(stringContent, -1)
	for _, match := range matches {
		newMul := Mul{}
		for i, val := range match {
			name := n1[i]
			if name == "firstNumber" {
				firstNumber, err := strconv.Atoi(val)
				if err != nil {
					fmt.Println(err)
				}
				newMul.firstNumber = firstNumber
			}
			if name == "secondNumber" {

				secondNumber, err := strconv.Atoi(val)
				if err != nil {
					fmt.Println(err)
				}
				newMul.secondNumber = secondNumber
			}

			validMuls = append(validMuls, newMul)
		}
	}
	return validMuls
}

func CalculateMulString(stringContent string) int {
	numbersString := strings.Replace(stringContent, "mul(", "", -1)
	numbersString = strings.Replace(numbersString, ")", "", -1)
	numbersSlice := strings.Split(numbersString, ",")
	firstNumber, err := strconv.Atoi(numbersSlice[0])
	if err != nil {
		fmt.Println(err)
	}
	secondNumber, err := strconv.Atoi(numbersSlice[1])
	if err != nil {
		fmt.Println(err)
	}
	return firstNumber * secondNumber
}

func ExtractDoString(stringContent string) []string {
	re := regexp.MustCompile(`do\(\)(?<do>.*?)don't\(\)`)
	n1 := re.SubexpNames()
	matches := re.FindAllStringSubmatch(stringContent, -1)

	doStrings := []string{}
	for _, match := range matches {
		for i, val := range match {
			name := n1[i]
			if name == "do" {
				doStrings = append(doStrings, val)
			}
		}
	}
	return doStrings
}

func ExtractStartAndEndingDoString(stringContent string) []string {
	re := regexp.MustCompile(`(?<start>.*?)don't\(\)(.*)do\(\)(?<end>.*)`)
	n1 := re.SubexpNames()
	matches := re.FindAllStringSubmatch(stringContent, -1)

	stringSlices := []string{}
	for _, match := range matches {
		for i, val := range match {
			name := n1[i]
			if name == "start" {
				stringSlices = append(stringSlices, val)
			}
			if name == "end" {
				stringSlices = append(stringSlices, val)
			}
		}
	}
	return stringSlices
}

func ExtractDoMuls(stringContent string) []Mul {
	validMuls := []Mul{}

	re := regexp.MustCompile(`do\((?P<firstNumber>[0-9]{1,3}),(?P<secondNumber>[0-9]{1,3})\)`)
	n1 := re.SubexpNames()
	matches := re.FindAllStringSubmatch(stringContent, -1)
	for _, match := range matches {
		newMul := Mul{}
		for i, val := range match {
			name := n1[i]
			if name == "firstNumber" {
				firstNumber, err := strconv.Atoi(val)
				if err != nil {
					fmt.Println(err)
				}
				newMul.firstNumber = firstNumber
			}
			if name == "secondNumber" {

				secondNumber, err := strconv.Atoi(val)
				if err != nil {
					fmt.Println(err)
				}
				newMul.secondNumber = secondNumber
			}

			validMuls = append(validMuls, newMul)
		}
	}
	return validMuls
}

func SumMuls(muls []Mul) int {
	total := 0
	for _, mul := range muls {
		multipliedMuls := mul.firstNumber * mul.secondNumber
		total += multipliedMuls
	}
	return total
}

func ExtractValidOperators(stringContent string) []string {
	re := regexp.MustCompile(`(?<mul>mul\([0-9]{1,3},[0-9]{1,3}\))|(?<do>do\(\))|(?<dont>don't\(\))`)
	n1 := re.SubexpNames()
	matches := re.FindAllStringSubmatch(stringContent, -1)
	validOperators := []string{}
	for _, match := range matches {
		for i, val := range match {
			name := n1[i]
			if name == "mul" {
				validOperators = append(validOperators, val)

			}
			if name == "do" {
				validOperators = append(validOperators, val)

			}
			if name == "dont" {
				validOperators = append(validOperators, val)

			}
		}
	}
	return validOperators
}

func WriteToFile(path string, content string) {

	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fmt.Fprint(file, content)
}
