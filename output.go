package main

import "fmt"

// Print - string ...

// PrintBlack ...
func PrintBlack(str string) {
	fmt.Printf("%s%d;40m%s%s", ESCAPECHAR, Black, str, ESCAPECHARCLOSE)
}

// PrintBlackln ...
func PrintBlackln(str string) {
	fmt.Printf("%s%d;40m%s%s\n", ESCAPECHAR, Black, str, ESCAPECHARCLOSE)
}

// PrintGreen ...
func PrintGreen(str string) {
	fmt.Printf("%s%d;40m%s%s", ESCAPECHAR, Green, str, ESCAPECHARCLOSE)
}

// PrintGreenln ...
func PrintGreenln(str string) {
	fmt.Printf("%s%d;40m%s%s\n", ESCAPECHAR, Green, str, ESCAPECHARCLOSE)
}

// PrintRed ...
func PrintRed(str string) {
	fmt.Printf("%s%d;40m%s%s", ESCAPECHAR, Red, str, ESCAPECHARCLOSE)
}

// PrintRedln ...
func PrintRedln(str string) {
	fmt.Printf("%s%d;40m%s%s\n", ESCAPECHAR, Red, str, ESCAPECHARCLOSE)
}

// PrintYellow ...
func PrintYellow(str string) {
	fmt.Printf("%s%d;40m%s%s", ESCAPECHAR, Yellow, str, ESCAPECHARCLOSE)
}

// PrintYellowln ...
func PrintYellowln(str string) {
	fmt.Printf("%s%d;40m%s%s\n", ESCAPECHAR, Yellow, str, ESCAPECHARCLOSE)
}

// PrintBlue ...
func PrintBlue(str string) {
	fmt.Printf("%s%d;40m%s%s", ESCAPECHAR, Blue, str, ESCAPECHARCLOSE)
}

// PrintBlueln ...
func PrintBlueln(str string) {
	fmt.Printf("%s%d;40m%s%s\n", ESCAPECHAR, Blue, str, ESCAPECHARCLOSE)
}

// PrintMagenta ...
func PrintMagenta(str string) {
	fmt.Printf("%s%d;40m%s%s", ESCAPECHAR, Magenta, str, ESCAPECHARCLOSE)
}

// PrintMagentaln ...
func PrintMagentaln(str string) {
	fmt.Printf("%s%d;40m%s%s\n", ESCAPECHAR, Magenta, str, ESCAPECHARCLOSE)
}

// PrintCyan ...
func PrintCyan(str string) {
	fmt.Printf("%s%d;40m%s%s", ESCAPECHAR, Cyan, str, ESCAPECHARCLOSE)
}

// PrintCyanln ...
func PrintCyanln(str string) {
	fmt.Printf("%s%d;40m%s%s\n", ESCAPECHAR, Cyan, str, ESCAPECHARCLOSE)
}

// PrintWhite ...
func PrintWhite(str string) {
	fmt.Printf("%s%d;40m%s%s", ESCAPECHAR, White, str, ESCAPECHARCLOSE)
}

// PrintWhiteln ...
func PrintWhiteln(str string) {
	fmt.Printf("%s%d;40m%s%s\n", ESCAPECHAR, White, str, ESCAPECHARCLOSE)
}

// PrintCustomRGB ...
func PrintCustomRGB(color RGB, str string) {
	fmt.Printf("%s38;2;%d;%d;%dm %s%s", ESCAPECHAR3, color.R, color.G, color.B, str, ESCAPECHARCLOSE)
}

// PrintCustomRGBln ...
func PrintCustomRGBln(color RGB, str string) {
	fmt.Printf("%s38;2;%d;%d;%dm %s%s\n", ESCAPECHAR3, color.R, color.G, color.B, str, ESCAPECHARCLOSE)
}

// PrintCustomHEX ...
func PrintCustomHEX(color HEX, str string) {
	c := color.ToRGB()
	fmt.Printf("%s38;2;%d;%d;%dm %s%s", ESCAPECHAR3, c.R, c.G, c.B, str, ESCAPECHARCLOSE)
}

// PrintCustomHEXln ...
func PrintCustomHEXln(color HEX, str string) {
	c := color.ToRGB()
	fmt.Printf("%s38;2;%d;%d;%dm %s%s\n", ESCAPECHAR3, c.R, c.G, c.B, str, ESCAPECHARCLOSE)
}
