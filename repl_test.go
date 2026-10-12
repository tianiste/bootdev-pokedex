package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " Hello world ",
			expected: []string{"Hello", "world"},
		},
		{
			input:    "Hey there",
			expected: []string{"Hey", "there"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("length of input != expected (len(%d) != len(%d))", len(actual), len(c.expected))
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("word doesent match expected word (%s != %s)", word, expectedWord)
			}
		}
	}
}
