package main

import "fmt"

// format - returns string ...

// FormatBlack ...
func FormatBlack(str string) string {
	return fmt.Sprintf("%s%d;1m%s%s0m", ESCAPECHAR, Black, str, ESCAPECHAR)
}

// FormatBlackln ...
func FormatBlackln(str string) string {
	return fmt.Sprintf("%s%d;1m%s%s0m\n", ESCAPECHAR, Black, str, ESCAPECHAR)
}

// FormatGreen ...
func FormatGreen(str string) string {
	return fmt.Sprintf("%s%d;1m%s%s0m", ESCAPECHAR, Green, str, ESCAPECHAR)
}

// FormatGreenln ...
func FormatGreenln(str string) string {
	return fmt.Sprintf("%s%d;1m%s%s0m\n", ESCAPECHAR, Green, str, ESCAPECHAR)
}

// FormatRed ...
func FormatRed(str string) string {
	return fmt.Sprintf("%s%d;1m%s%s0m", ESCAPECHAR, Red, str, ESCAPECHAR)
}

// FormatRedln ...
func FormatRedln(str string) string {
	return fmt.Sprintf("%s%d;1m%s%s0m\n", ESCAPECHAR, Red, str, ESCAPECHAR)
}

// FormatYellow ...
func FormatYellow(str string) string {
	return fmt.Sprintf("%s%d;1m%s%s0m", ESCAPECHAR, Yellow, str, ESCAPECHAR)
}

// FormatYellowln ...
func FormatYellowln(str string) string {
	return fmt.Sprintf("%s%d;1m%s%s0m\n", ESCAPECHAR, Yellow, str, ESCAPECHAR)
}

// FormatBlue ...
func FormatBlue(str string) string {
	return fmt.Sprintf("%s%d;1m%s%s0m", ESCAPECHAR, Blue, str, ESCAPECHAR)
}

// FormatBlueln ...
func FormatBlueln(str string) string {
	return fmt.Sprintf("%s%d;1m%s%s0m\n", ESCAPECHAR, Blue, str, ESCAPECHAR)
}

// FormatMagenta ...
func FormatMagenta(str string) string {
	return fmt.Sprintf("%s%d;1m%s%s0m", ESCAPECHAR, Magenta, str, ESCAPECHAR)
}

// FormatMagentaln ...
func FormatMagentaln(str string) string {
	return fmt.Sprintf("%s%d;1m%s%s0m\n", ESCAPECHAR, Magenta, str, ESCAPECHAR)
}

// FormatCyan ...
func FormatCyan(str string) string {
	return fmt.Sprintf("%s%d;1m%s%s0m", ESCAPECHAR, Cyan, str, ESCAPECHAR)
}

// FormatCyanln ...
func FormatCyanln(str string) string {
	return fmt.Sprintf("%s%d;1m%s%s0m\n", ESCAPECHAR, Cyan, str, ESCAPECHAR)
}

// FormatWhite ...
func FormatWhite(str string) string {
	return fmt.Sprintf("%s%d;1m%s%s0m", ESCAPECHAR, White, str, ESCAPECHAR)
}

// FormatWhiteln ...
func FormatWhiteln(str string) string {
	return fmt.Sprintf("%s%d;1m%s%s0m\n", ESCAPECHAR, White, str, ESCAPECHAR)
}

// TODO Add escape close args

// FormatCustom ...
func FormatCustom(color RGB, str string) string {
	return fmt.Sprintf("%s38;2;%d;%d;%dm %s%s\n", ESCAPECHAR3, color.R, color.G, color.B, str, ESCAPECHARCLOSE)
}

// FormatCustomln ...
func FormatCustomln(color RGB, str string) string {
	return fmt.Sprintf("%s38;2;%d;%d;%dm %s%s\n", ESCAPECHAR3, color.R, color.G, color.B, str, ESCAPECHARCLOSE)
}
