package main

import "testing"

func TestSubstr(t *testing.T) {
	tests := []struct {
		s   string
		ans int
	}{
		{"abcadcbb", 4},
		//{"pwwkew", 3},
		//
		//{"", 0},
		//{"bbbbbb", 1},
		//{"abcabcabcd", 4},
	}

	for _, tt := range tests {
		actual := lengthOfNonRepeatingSubStr(tt.s)
		if actual != tt.ans {
			t.Errorf("Got %d for input %s;"+"excepted %d", actual, tt.s, tt.ans)
		}
	}
}

func BenchmarkSubStr(b *testing.B) {
	s := "pwwkew"
	ans := 3
	for i := 0; i < b.N; i++ {
		actual := lengthOfNonRepeatingSubStr(s)
		if actual != ans {
			b.Errorf("Got %d for input %s;"+"excepted %d", actual, s, ans)
		}
	}

}
