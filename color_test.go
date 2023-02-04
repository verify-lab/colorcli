package colorcli

import (
	"testing"
)

func TestColorize(t *testing.T) {
	tests := []struct {
		name  string
		color string
		txt   any
		want  string
	}{
		// echo -e "\033[35mtest\033[0m"
		{"no color", "", "test", "test"},
		{"black", Black, "test", "\033[30mtest\033[0m"},
		{"red", Red, "test", "\033[31mtest\033[0m"},
		{"green", Green, "test", "\033[32mtest\033[0m"},
		{"yellow", Yellow, "test", "\033[33mtest\033[0m"},
		{"blue", Blue, "test", "\033[34mtest\033[0m"},
		{"purple", Purple, "test", "\033[35mtest\033[0m"},
		{"cyan", Cyan, "test", "\033[36mtest\033[0m"},
		{"gray", Gray, "test", "\033[37mtest\033[0m"},
		{"white", White, "test", "\033[97mtest\033[0m"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Colorize(tt.color, tt.txt); got != tt.want {
				t.Errorf("Colorize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestColorizeWithBackgound(t *testing.T) {
	tests := []struct {
		name     string
		txtColor string
		bgColor  string
		txt      any
		want     string
	}{
		// echo -e "\033[31m\033[46mtext\033[0m"
		{"no color", "", "", "text", "text"},
		{"red over cyan", Red, CyanBackground, "text", "\033[31m\033[46mtext\033[0m"},
		{"green over white", Green, WhiteBackground, "text", "\033[32m\033[107mtext\033[0m"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ColorizeWithBackgound(tt.txtColor, tt.bgColor, tt.txt); got != tt.want {
				t.Errorf("ColorizeWithBackgound() = %v, want %v", got, tt.want)
			}
		})
	}
}
