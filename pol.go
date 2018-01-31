package calc

import (
	"github.com/LepikovStan/calc/stack"
	"strconv"
	"strings"
)

var allowedOperands = map[string]map[string]int{
	"^": map[string]int{"priority": 10},
	"*": map[string]int{"priority": 7},
	"/": map[string]int{"priority": 7},
	"+": map[string]int{"priority": 5},
	"-": map[string]int{"priority": 5},
	"(": map[string]int{"priority": 2},
	")": map[string]int{"priority": 2},
}

func isAllowedOperand(char string) bool {
	_, ok := allowedOperands[char]
	return ok
}

func splitExpr(expr string) []string {
	items := strings.Split(expr, "")
	for _, item := range items {
		item = strings.TrimSpace(item)
	}
	return items
}

/*
 *  Method translate infix notation to reverse poland
 */
func GetPoland(expr string) *stack.Stack {
	// no output string, stack resultStk instead
	resultStk := stack.New()
	operandStk := stack.New()
	// it's need to compose number from several chars, declare a variable for future number
	number := ""

	for _, char := range splitExpr(expr) {
		if char == " " {
			continue
		}
		_, errFloat := strconv.ParseFloat(char, 64)

		// if char is not float and it's allowed operand
		if errFloat != nil && isAllowedOperand(char) {
			if number != "" {
				resultStk.Push(number)
			}
			number = ""

			// if char is `(` push it to operand stack
			if char == "(" {
				operandStk.Push(char)
				continue
			}
			if char == ")" {
				// while char from operand stack is not '(' pop operands and push their to result stack
				for {
					headElem, err := operandStk.Pop()
					if err != nil {
						panic(err)
					}
					if headElem != "(" {
						resultStk.Push(headElem)
					} else {
						break
					}
				}
				continue
			}

			// if char is operand and it's priority less or equal head operand stack element
			// then push head of operand stack to result stack
			if allowedOperands[char]["priority"] <= allowedOperands[operandStk.Head()]["priority"] {
				headElem, err := operandStk.Pop()
				if err != nil {
					panic(err)
				}
				resultStk.Push(headElem)
			}

			// push char to operand stack
			operandStk.Push(char)
		} else {
			// compose number
			number += char
		}
	}

	// push last number in infix string to result stack
	if number != "" {
		resultStk.Push(number)
	}

	// move all operands from operand stack to result stack
	for {
		operand, err := operandStk.Pop()
		if err != nil {
			break
		}
		resultStk.Push(operand)
	}
	resultStk.Reverse()
	return resultStk
}

/*
 * Method range all elements from result stack and join it to string
 * it's need for test only
 * for me it's more cosily to test strings
 */
func GetPolandString(expr string) string {
	result := ""
	stk := GetPoland(expr)
	for {
		item, err := stk.Pop()
		if err != nil {
			break
		}
		result += item + " "
	}
	return result
}
