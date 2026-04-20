package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCleanInput(t *testing.T) {
	tests := map[string]struct {
		input string
		want  []string
	}{
		"all_spaces":     {input: "   ", want: []string{}},
		"all_lower":      {input: "hello world", want: []string{"hello", "world"}},
		"all_upper":      {input: "HELLO WORLD", want: []string{"hello", "world"}},
		"trailing_space": {input: "hello world  ", want: []string{"hello", "world"}},
		"leading_space":  {input: "  hello world", want: []string{"hello", "world"}},
		"extra_spaces":   {input: "hello   world", want: []string{"hello", "world"}},
		"complex":        {input: "  No pAin  no gain  ", want: []string{"no", "pain", "no", "gain"}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := CleanInput(tc.input)
			diff := cmp.Diff(tc.want, got)
			if diff != "" {
				t.Errorf("CleanInput() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
