package tests

import (
	"testing"

	"github.com/TommyFiga/go-mathy/calc"
)


func TestCalculator(t *testing.T) {
	tests := []struct {
		a,b 		float64
		op 			string
		expected	float64
		hasErr		bool
	}{
		{2, 3, "+", 5, false},
		{10, 4, "-", 6, false},
		{5, 2, "*", 10, false},
		{3, 3, "x", 9, false},
		{3, 3, "X", 9, false},
		{9, 3, "/", 3, false},
		{1, 0, "/", 0, true},
		{2, 2, "%", 0, true},
	}

	for _, tc := range tests {
		res, err := calc.Calculator(tc.a, tc.op, tc.b)

		if tc.hasErr && err == nil {
			t.Errorf("expected error for %v %s %v, got nil", tc.a, tc.op, tc.b)
		}

		if !tc.hasErr && err != nil {
			t.Errorf("unexpected error for %v %s %v: %v", tc.a, tc.op, tc.b, err)
		}

		if !tc.hasErr && res != tc.expected {
			t.Errorf("expected %v, got %v for %v %s %v", tc.expected, res, tc.a, tc.op, tc.b)
		}
	}
}


func TestConvertToNumber(t *testing.T) {
	tests := []struct{
		input		string
		expected	float64
		hasErr		bool
	}{
		{"-5", -5, false},
		{"3.14", 3.14, false},
		{"3", 3, false},
		{"", 0, true},
		{"",0, true},
	}

	for _, tc := range tests {
		res, err := calc.ConvertToNumber(tc.input)

		if tc.hasErr && err == nil {
			t.Errorf("expected error for %s, got nil", tc.input)
		}
		if !tc.hasErr && err != nil {
			t.Errorf("unexpected error for %s: %v", tc.input, err)
		}
		if !tc.hasErr && res != tc.expected {
			t.Errorf("expected %v, got %v for %s", tc.expected, res, tc.input)
		}
	} 
}