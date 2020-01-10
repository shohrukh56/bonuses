package main

import "testing"

func Test_bonuses(t *testing.T) {
	tests := []struct {
		name string
		sales []int
		want int
	}{
		// TODO: Add test cases.
		{"all bonuses lower than salesborder", []int{1, 5, 10, 15}, 0 },
		{"all bonuses higer than salesborder", []int{11_000, 15_000, 13_000, 20_000}, 950},
		{"equal to salesborder", []int{10_000,10_000,10_000,10_000,}, 0 },
		{"mixed nubmer", []int{10_000, 15_000, 9_000, 18_000}, 650},
	}
	for _, tt := range tests {
			if got := bonuses(tt.sales); got != tt.want {
				t.Error("bonuseAmount test", tt.name,"got", got, "want",tt.want )
			}
		}
	}
