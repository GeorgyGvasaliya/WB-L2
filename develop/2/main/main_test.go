package main

import "testing"

func TestFind(t *testing.T) {
	testTable := []struct {
		have     string
		expected string
	}{
		{
			have:     "a4bc2d5e",
			expected: "aaaabccddddde",
		},
		{
			have:     "abcd",
			expected: "abcd",
		},
		{
			have:     "",
			expected: "",
		},
		{
			have:     `qwe\4\5`,
			expected: "qwe45",
		},
		{
			have:     `qwe\45`,
			expected: "qwe44444",
		},
		{
			have:     `qwe\\5`,
			expected: `qwe\\\\\`,
		},
	}

	for _, tc := range testTable {
		got := Unpack(tc.have)
		if got != tc.expected {
			t.Errorf("Expected %s, got %s", tc.expected, got)
		}
	}
}
