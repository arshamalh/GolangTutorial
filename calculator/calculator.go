package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type OperandError rune

func (oe OperandError) Error() string {
	return fmt.Sprintf("operand %s is not valid", string(oe))
}

type Process struct {
	Number1 float64
	Number2 float64
	Operand rune
}

func main() {
	for {
		// Get input
		input := getInput()

		// parse input
		process, err := parse(input)
		if err != nil {
			fmt.Println(err)
			continue
		}

		// Calculate
		result, err := process.calculate()
		if err != nil {
			fmt.Println(err)
			continue
		}

		// Print result
		fmt.Println(result)
	}
}

func getInput() string {
	var input string
	fmt.Print("> ")
	fmt.Scanln(&input)
	return input
}

func parse(input string) (*Process, error) {
	for _, operand := range "+-*/%^" {
		subs := strings.Split(input, string(operand))
		if len(subs) != 1 {
			number_1, err := strconv.ParseFloat(subs[0], 64)
			if err != nil {
				return nil, err
			}
			number_2, err := strconv.ParseFloat(subs[1], 64)
			if err != nil {
				return nil, err
			}
			return &Process{
				Number1: number_1,
				Number2: number_2,
				Operand: operand,
			}, nil
		}
	}
	return nil, OperandError(' ')
}

func (p *Process) calculate() (float64, error) {
	switch p.Operand {
	case '+':
		return p.Number1 + p.Number2, nil
	case '-':
		return p.Number1 - p.Number2, nil
	case '*':
		return p.Number1 * p.Number2, nil
	case '/':
		if p.Number2 == 0 {
			return 0, fmt.Errorf("dividing by 0 is not possible")
		}
		return p.Number1 / p.Number2, nil
	case '^':
		return math.Pow(p.Number1, p.Number2), nil
	case '%':
		return math.Mod(p.Number1, p.Number2), nil
	}
	return 0, OperandError(p.Operand)
}
