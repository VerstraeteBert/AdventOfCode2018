package main

import (
	"testing"
)

func BenchmarkConcurrent(b *testing.B) {
	for n := 0; n < b.N; n++ {
		concurrent()
	}
}

func BenchmarkNonConcurrent(b *testing.B) {
	for n := 0; n < b.N; n++ {
		nonConcurrent()
	}
}

func Test_getDupAndTripLetters(t *testing.T) {
	tests := []struct {
		name         string
		testString   string
		wantGotTwo   bool
		wantGotThree bool
	}{
		{"No Duplicates", "azerty", false, false },
		{"One double occurence", "azearty", true, false},
		{"One triple occurence", "azerrtyr", false, true},
		{"Both double and triple", "azerttytz", true, true},
		{"Multiple double occurence", "azzertye", true, false},
		{"Multiple triple occurence", "azzzertyyy", false, true},
		{"No double or triple", "azzzzzzzertyyyyyyyyyy", false, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotGotTwo, gotGotThree := getDupAndTripLetters(tt.testString)
			if gotGotTwo != tt.wantGotTwo {
				t.Errorf("getDupAndTripLetters() gotGotTwo = %v, want %v", gotGotTwo, tt.wantGotTwo)
			}
			if gotGotThree != tt.wantGotThree {
				t.Errorf("getDupAndTripLetters() gotGotThree = %v, want %v", gotGotThree, tt.wantGotThree)
			}
		})
	}
}
