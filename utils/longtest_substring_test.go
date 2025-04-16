package utils

import "testing"

func TestLongestUniqueSubstring(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"abcabcbb", "abc"},
		{"bbbbb", "b"},
		{"pwwkew", "wke"},
		{"", ""},
		{"abcdef", "abcdef"},
		{"aab", "ab"},
		{"dvdf", "vdf"},
	}

	for _, test := range tests {
		got := LongestSubstring(test.input)
		if test.expected != got {
			t.Errorf("Expected: %#v, got: %#v", test.expected, got)
		}
	}
}
