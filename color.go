package colorcli

import (
	"fmt"
	"strings"
)

// https://en.wikipedia.org/wiki/ANSI_escape_code#Colors
// https://jeffkreeftmeijer.com/vim-16-color

const (
	// Special
	Reset     = "\033[0m"
	Bold      = "\033[1m"
	Underline = "\033[4m"
	// Text colors
	Black  = "\033[30m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	Gray   = "\033[37m"
	White  = "\033[97m"
	// Background colors
	BlackBackground  = "\033[40m"
	RedBackground    = "\033[41m"
	GreenBackground  = "\033[42m"
	YellowBackground = "\033[43m"
	BlueBackground   = "\033[44m"
	PurpleBackground = "\033[45m"
	CyanBackground   = "\033[46m"
	GrayBackground   = "\033[47m"
	WhiteBackground  = "\033[107m"
)

func Colorize(color string, s any) string {
	var sb strings.Builder

	sb.WriteString(color)

	switch v := s.(type) {
	case string:
		sb.WriteString(v)
	default:
		sb.WriteString(fmt.Sprint(v))
	}

	if color != "" {
		sb.WriteString(Reset)
	}

	return sb.String()
}

func ColorizeWithBackgound(txtColor, bgColor string, s any) string {
	var sb strings.Builder

	sb.WriteString(txtColor)
	sb.WriteString(bgColor)

	switch v := s.(type) {
	case string:
		sb.WriteString(v)
	default:
		sb.WriteString(fmt.Sprint(v))
	}

	if txtColor != "" && bgColor != "" {
		sb.WriteString(Reset)
	}

	return sb.String()
}
