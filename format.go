package niji

import "fmt"

// format - returns string ...

// FormatBlack ...
func FormatBlack(str string) string {
	return fmt.Sprintf("%s%d;1m%s%s", ESCAPECHAR, Black, str, ESCAPECHARCLOSE)
}

// FormatBlackln ...
func FormatBlackln(str string) string {
	return fmt.Sprintf("%s%d;1m%s%s\n", ESCAPECHAR, Black, str, ESCAPECHARCLOSE)
}

// FormatGreen ...
func FormatGreen(str string) string {
	return fmt.Sprintf("%s%d;1m%s%s", ESCAPECHAR, Green, str, ESCAPECHARCLOSE)
}

// FormatGreenln ...
func FormatGreenln(str string) string {
	return fmt.Sprintf("%s%d;1m%s%s\n", ESCAPECHAR, Green, str, ESCAPECHARCLOSE)
}

// FormatRed ...
func FormatRed(str string) string {
	return fmt.Sprintf("%s%d;1m%s%s", ESCAPECHAR, Red, str, ESCAPECHARCLOSE)
}

// FormatRedln ...
func FormatRedln(str string) string {
	return fmt.Sprintf("%s%d;1m%s%s\n", ESCAPECHAR, Red, str, ESCAPECHARCLOSE)
}

// FormatYellow ...
func FormatYellow(str string) string {
	return fmt.Sprintf("%s%d;1m%s%s", ESCAPECHAR, Yellow, str, ESCAPECHARCLOSE)
}

// FormatYellowln ...
func FormatYellowln(str string) string {
	return fmt.Sprintf("%s%d;1m%s%s\n", ESCAPECHAR, Yellow, str, ESCAPECHARCLOSE)
}

// FormatBlue ...
func FormatBlue(str string) string {
	return fmt.Sprintf("%s%d;1m%s%s", ESCAPECHAR, Blue, str, ESCAPECHARCLOSE)
}

// FormatBlueln ...
func FormatBlueln(str string) string {
	return fmt.Sprintf("%s%d;1m%s%s\n", ESCAPECHAR, Blue, str, ESCAPECHARCLOSE)
}

// FormatMagenta ...
func FormatMagenta(str string) string {
	return fmt.Sprintf("%s%d;1m%s%s", ESCAPECHAR, Magenta, str, ESCAPECHARCLOSE)
}

// FormatMagentaln ...
func FormatMagentaln(str string) string {
	return fmt.Sprintf("%s%d;1m%s%s\n", ESCAPECHAR, Magenta, str, ESCAPECHARCLOSE)
}

// FormatCyan ...
func FormatCyan(str string) string {
	return fmt.Sprintf("%s%d;1m%s%s", ESCAPECHAR, Cyan, str, ESCAPECHARCLOSE)
}

// FormatCyanln ...
func FormatCyanln(str string) string {
	return fmt.Sprintf("%s%d;1m%s%s\n", ESCAPECHAR, Cyan, str, ESCAPECHARCLOSE)
}

// FormatWhite ...
func FormatWhite(str string) string {
	return fmt.Sprintf("%s%d;1m%s%s", ESCAPECHAR, White, str, ESCAPECHARCLOSE)
}

// FormatWhiteln ...
func FormatWhiteln(str string) string {
	return fmt.Sprintf("%s%d;1m%s%s\n", ESCAPECHAR, White, str, ESCAPECHARCLOSE)
}

// FormatCustomRGB ...
func FormatCustomRGB(color RGB, str string) string {
	return fmt.Sprintf("%s38;2;%d;%d;%dm %s%s", ESCAPECHAR3, color.R, color.G, color.B, str, ESCAPECHARCLOSE)
}

// FormatCustomRGBln ...
func FormatCustomRGBln(color RGB, str string) string {
	return fmt.Sprintf("%s38;2;%d;%d;%dm %s%s\n", ESCAPECHAR3, color.R, color.G, color.B, str, ESCAPECHARCLOSE)
}

// FormatCustomHEX ...
func FormatCustomHEX(color HEX, str string) string {
	c := color.ToRGB()
	return fmt.Sprintf("%s38;2;%d;%d;%dm %s%s", ESCAPECHAR3, c.R, c.G, c.B, str, ESCAPECHARCLOSE)
}

// FormatCustomHEXln ...
func FormatCustomHEXln(color HEX, str string) string {
	c := color.ToRGB()
	return fmt.Sprintf("%s38;2;%d;%d;%dm %s%s\n", ESCAPECHAR3, c.R, c.G, c.B, str, ESCAPECHARCLOSE)
}
