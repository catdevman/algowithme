package validparentheses

import (
	"fmt"
	"testing"
)

func TestValidParentheses(t *testing.T) {
	tests := map[string]struct {
		input    string
		expected bool
	}{
		"single": {
			"()",
			true,
		},
		"multiple pairs": {
			"()[]{}",
			true,
		},
		"mismatch": {
			")[",
			false,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			out := ValidParentheses(test.input)
			if out != test.expected {
				t.Fatal(fmt.Sprintf("expected %t but got %t", test.expected, out))
			}
		})
	}
}
