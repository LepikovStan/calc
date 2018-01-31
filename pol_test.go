package calc

import (
	"testing"
)

func TestGetPolandString(t *testing.T) {
	results := map[string]string{
		"1 + 2":                 "1 2 + ",
		"1.5 + 3.2":             "1.5 3.2 + ",
		"3 + 4 * 2":             "3 4 2 * + ",
		"15 + 23":               "15 23 + ",
		"(1+2)*4+3":             "1 2 + 4 * 3 + ",
		"3 + 4 * 2 / (1 - 5)^2": "3 4 2 * 1 5 - 2 ^ / + ",
	}

	for infix, poland := range results {
		res := GetPolandString(infix)
		t.Logf("Expexted `%s`, result `%s`", poland, res)
		if res != poland {
			t.Fatalf("Expexted `%s`, result `%s`", poland, res)
		}
	}
}
