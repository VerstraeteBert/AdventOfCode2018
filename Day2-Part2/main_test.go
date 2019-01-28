package main

import "testing"

func Test_compare(t *testing.T) {
	type args struct {
		wordOne string
		wordTwo string
	}
	tests := []struct {
		name        string
		args        args
		wantCommon  string
		wantIsClose bool
	}{
		{"No common letters", args{"azerty", "qsdfgh"}, "", false},
		{"Two away", args{"azerty", "bzertu"}, "", false},
		{"One away", args{"azerty", "azertu"}, "azert", true},
		{"Empty", args{"", ""}, "", false},
		{"Different length", args{"a", "aa"}, "", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCommon, gotIsClose := compare(tt.args.wordOne, tt.args.wordTwo)
			if gotCommon != tt.wantCommon {
				t.Errorf("compare() gotCommon = %v, want %v", gotCommon, tt.wantCommon)
			}
			if gotIsClose != tt.wantIsClose {
				t.Errorf("compare() gotIsClose = %v, want %v", gotIsClose, tt.wantIsClose)
			}
		})
	}
}
