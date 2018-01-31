package calc

import (
	"testing"
)

func TestCalcPoland(t *testing.T) {
	results := map[string]float64{
		"1 + 2":                 3,
		"1.5 + 3.2":             4.7,
		"3 + (4 * 2)":           11,
		"(1+2)*4+3":             15,
		"15 + 23":               38,
		"3 + 4 * 2 / (1 - 5)^2": 3.5,
	}

	for infix, expected := range results {
		poland := GetPoland(infix)
		result := CalcPoland(poland)
		if expected != result {
			t.Fatalf("Expexted `%s=%f`, result `%f`", infix, expected, result)
		}
	}
}
