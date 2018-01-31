package calc

import (
	"github.com/LepikovStan/calc/stack"
	"math"
	"strconv"
)

var allowedOperations = map[string]func(float64, float64) float64{
	"^": func(x float64, y float64) float64 {
		return math.Pow(x, y)
	},
	"*": func(x, y float64) float64 {
		return x * y
	},
	"/": func(x, y float64) float64 {
		return x / y
	},
	"+": func(x, y float64) float64 {
		return x + y
	},
	"-": func(x, y float64) float64 {
		return x - y
	},
}

/*
 * Method get result stack in poland notation and calculate it
 */
func CalcPoland(stk *stack.Stack) float64 {
	resultStack := stack.New()
	for {
		elem, err := stk.Pop()

		if err != nil {
			break
		}

		_, err = strconv.ParseFloat(elem, 64)

		if err != nil {
			second, _ := resultStack.Pop()
			first, _ := resultStack.Pop()
			x, _ := strconv.ParseFloat(first, 64)
			y, _ := strconv.ParseFloat(second, 64)
			res := allowedOperations[elem](x, y)
			resultStack.Push(strconv.FormatFloat(res, 'f', -1, 64))
		} else {
			resultStack.Push(elem)
		}
	}
	res, err := resultStack.Pop()
	if err != nil {
		panic(err)
	}
	r, err := strconv.ParseFloat(res, 64)
	if err != nil {
		panic(err)
	}
	return r
}
