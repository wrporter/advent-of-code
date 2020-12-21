package pda

import (
	"fmt"
	"testing"
)

func TestPDA_Match(t *testing.T) {
	type inputResults struct {
		input string
		want  bool
	}
	tests := []struct {
		startSymbol  string
		rules        []string
		inputResults []inputResults
	}{
		{
			startSymbol: "S",
			rules: []string{
				"S: a S b",
				"S: !",
			},
			inputResults: []inputResults{
				{input: "ab", want: true},
				{input: "aabb", want: true},
				{input: "!", want: true},
				{input: "aaaabbbb", want: true},
				{input: "ababa", want: false},
				{input: "bbaaa", want: false},
			},
		},
		{
			startSymbol: "S",
			rules: []string{
				"S: b A b",
				"S: a A a",
				"S: a",
				"S: b",
				"A: b A b",
				"A: a A a",
				"A: a",
				"A: b",
				"A: b A a",
				"A: a A b",
				"A: !",
			},
			inputResults: []inputResults{
				{input: "a", want: true},
				{input: "b", want: true},
				{input: "baaababababababbbbbbab", want: true},
				{input: "abbbbbabbbabbbbbbbabba", want: true},
				{input: "babbab", want: true},
				{input: "baaaaba", want: false},
			},
		},
		{
			startSymbol: "S",
			rules: []string{
				"S: a_ N _tells_a_ N _ O",
				"N: boy",
				"N: girl",
				"O: a_story",
				"O: that_ S",
			},
			inputResults: []inputResults{
				{input: "a_boy_tells_a_girl_a_story", want: true},
				{input: "a_boy_tells_a_girl_that_a_boy_tells_a_girl_a_story", want: true},
				{input: "a_girl_tells_a_boy_a_story", want: true},
				{input: "a_girl_tells_a_boy_that_a_boy_tells_a_girl_that_a_girl_tells_a_boy_a_story", want: true},
				{input: "a_story_tells_a_boy_a_story", want: false},
				{input: "a_boy_tells_a_tells_a_boy_a_story", want: false},
			},
		},
		{
			startSymbol: "S",
			rules: []string{
				"S: a S a",
				"S: b S b",
				"S: !",
			},
			inputResults: []inputResults{
				{input: "abba", want: true},
			},
		},
		{
			startSymbol: "S",
			rules: []string{
				"S: U",
				"S: V",
				"U: T a U",
				"U: T a T",
				"V: T b V",
				"V: T b T",
				"T: a T b T",
				"T: b T a T",
				"T: !",
			},
			inputResults: []inputResults{
				{input: "aab", want: true},
				{input: "bbbbbbbbbbbbbbbbbbbbbbbbbbba", want: true},
				{input: "ababababababb", want: true},
				{input: "aaaaaaaaaabbbbbbbbbbb", want: true},
				{input: "aaaaaaaaaabbbbbbbbbb", want: false},
				{input: "ab", want: false},
			},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test_%d", i), func(t *testing.T) {
			p := &PDA{
				startSymbol: tt.startSymbol,
			}
			p.AddRules(tt.rules)

			for _, inputResult := range tt.inputResults {
				if got := p.Match(inputResult.input); got != inputResult.want {
					t.Errorf("Match() = %v, want %v - for %v", got, inputResult.want, inputResult.input)
				}
			}
		})
	}
}
