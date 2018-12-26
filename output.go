package main

import "fmt"

// Print - string ...

// PrintBlack ...
func PrintBlack(str string) {
	fmt.Printf("%s%d;40m%s%s0m", ESCAPECHAR, Black, str, ESCAPECHAR)
}

// PrintBlackln ...
func PrintBlackln(str string) {
	fmt.Printf("%s%d;40m%s%s0m\n", ESCAPECHAR, Black, str, ESCAPECHAR)
}

// PrintGreen ...
func PrintGreen(str string) {
	fmt.Printf("%s%d;40m%s%s0m", ESCAPECHAR, Green, str, ESCAPECHAR)
}

// PrintGreenln ...
func PrintGreenln(str string) {
	fmt.Printf("%s%d;40m%s%s0m\n", ESCAPECHAR, Green, str, ESCAPECHAR)
}

// PrintRed ...
func PrintRed(str string) {
	fmt.Printf("%s%d;40m%s%s0m", ESCAPECHAR, Red, str, ESCAPECHAR)
}

// PrintRedln ...
func PrintRedln(str string) {
	fmt.Printf("%s%d;40m%s%s0m\n", ESCAPECHAR, Red, str, ESCAPECHAR)
}

// PrintYellow ...
func PrintYellow(str string) {
	fmt.Printf("%s%d;40m%s%s0m", ESCAPECHAR, Yellow, str, ESCAPECHAR)
}

// PrintYellowln ...
func PrintYellowln(str string) {
	fmt.Printf("%s%d;40m%s%s0m\n", ESCAPECHAR, Yellow, str, ESCAPECHAR)
}

// PrintBlue ...
func PrintBlue(str string) {
	fmt.Printf("%s%d;40m%s%s0m", ESCAPECHAR, Blue, str, ESCAPECHAR)
}

// PrintBlueln ...
func PrintBlueln(str string) {
	fmt.Printf("%s%d;40m%s%s0m\n", ESCAPECHAR, Blue, str, ESCAPECHAR)
}

// PrintMagenta ...
func PrintMagenta(str string) {
	fmt.Printf("%s%d;40m%s%s0m", ESCAPECHAR, Magenta, str, ESCAPECHAR)
}

// PrintMagentaln ...
func PrintMagentaln(str string) {
	fmt.Printf("%s%d;40m%s%s0m\n", ESCAPECHAR, Magenta, str, ESCAPECHAR)
}

// PrintCyan ...
func PrintCyan(str string) {
	fmt.Printf("%s%d;40m%s%s0m", ESCAPECHAR, Cyan, str, ESCAPECHAR)
}

// PrintCyanln ...
func PrintCyanln(str string) {
	fmt.Printf("%s%d;40m%s%s0m\n", ESCAPECHAR, Cyan, str, ESCAPECHAR)
}

// PrintWhite ...
func PrintWhite(str string) {
	fmt.Printf("%s%d;40m%s%s0m", ESCAPECHAR, White, str, ESCAPECHAR)
}

// PrintWhiteln ...
func PrintWhiteln(str string) {
	fmt.Printf("%s%d;40m%s%s0m\n", ESCAPECHAR, White, str, ESCAPECHAR)
}

// PrintCustom ...
func PrintCustom(color RGB, str string) {
	fmt.Printf("%s38;2;%d;%d;%dm %s%s", ESCAPECHAR3, color.R, color.G, color.B, str, ESCAPECHARCLOSE)
}

// PrintCustomln ...
func PrintCustomln(color RGB, str string) {
	fmt.Printf("%s38;2;%d;%d;%dm %s%s\n", ESCAPECHAR3, color.R, color.G, color.B, str, ESCAPECHARCLOSE)
}
