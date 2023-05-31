package domain

import (
	"testing"
)

type Test struct {
	Input    string
	Expected []int
	Err      bool
}

// Example Simple Test
func TestStringToIdsSimple(t *testing.T) {
	out, err := stringToIds("-1")
	if err != nil {
		t.Error("err should be nil")
	}
	if len(out) != 1 {
		t.Error("length should be 1")
	} else if out[0] != -1 {
		t.Error("should be -1")
	}
}

// Example Table Driven Test
func TestStringToIdsTable(t *testing.T) {
	var tests []Test = []Test{{
		Input:    "1,2,3",
		Expected: []int{1, 2, 3},
		Err:      false,
	}, {
		Input:    "",
		Expected: []int{},
		Err:      true,
	}, {
		Input:    "1",
		Expected: []int{1},
		Err:      false,
	}, {
		Input:    "one,two,three",
		Expected: nil,
		Err:      true,
	}}

	for index, test := range tests {
		out, err := stringToIds(test.Input)
		if len(out) != len(test.Expected) {
			t.Errorf("%d != %d", len(out), len(test.Expected))
		} else {
			for i := range out {
				if out[i] != test.Expected[i] {
					t.Errorf("test #%d: expected %d, got %d", index+1, test.Expected[i], out[i])
				}
			}
		}
		hasErr := err != nil
		if hasErr && !test.Err {
			t.Errorf("expected no error, got %v", err)
		} else if !hasErr && test.Err {
			t.Errorf("expected error, got nil")
		}
	}
}
