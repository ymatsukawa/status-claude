package decorator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestColorize(t *testing.T) {
	tests := []struct {
		text  string
		color string
		want  string
	}{
		{"Hello", RED, "\033[31mHello\033[0m"},
		{"World", GREEN, "\033[32mWorld\033[0m"},
		{"Test", YELLOW, "\033[33mTest\033[0m"},
		{"Example", BLUE, "\033[34mExample\033[0m"},
		{"Default", DEFAULT, "Default"},
	}

	for _, tt := range tests {
		t.Run(tt.text+"_"+tt.color, func(t *testing.T) {
			got := Colorize(tt.text, tt.color)
			assert.Equal(t, tt.want, got)
		})
	}
}
