package main

import (
	"fmt"
	"testing"
)

func Test_calculateWrappingPaper(t *testing.T) {
	tests := []struct {
		presents   []string
		wantPaper  int
		wantRibbon int
	}{
		{[]string{"2x3x4"}, 58, 34},
		{[]string{"1x1x10"}, 43, 14},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test_%d", i), func(t *testing.T) {
			gotPaper, gotRibbon := calculate(tt.presents)
			if gotPaper != tt.wantPaper {
				t.Errorf("calculate() = %v, want paper %v", gotPaper, tt.wantPaper)
			}
			if gotRibbon != tt.wantRibbon {
				t.Errorf("calculate() = %v, want ribbon %v", gotRibbon, tt.wantRibbon)
			}
		})
	}
}
