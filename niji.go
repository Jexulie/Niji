package main

/*
Niji - 2018
* Does Work in Bash & ConEmu & Git-Bash
! Does not Work in Powershell & Command Prompt

*/

// ESCAPECHAR Escape chars ver 1
var ESCAPECHAR = "\u001b["

// ESCAPECHARCLOSE Escape char close 1
var ESCAPECHARCLOSE = "u001b[0m"

// ESCAPECHAR2 Escape chars ver 2
var ESCAPECHAR2 = "\x1b["

// ESCAPECHARCLOSE2 Escape char close 2
var ESCAPECHARCLOSE2 = "\x1b[0m"

// ESCAPECHAR3 Escape chars ver 3
var ESCAPECHAR3 = "\033["

// ESCAPECHARCLOSE3 Escape char close 3
var ESCAPECHARCLOSE3 = "\033[0m"

// DefaultColors type
type DefaultColors int

// DefaultColors Enum
const (
	Black DefaultColors = iota + 30
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
)

// RGB struct
type RGB struct {
	R int
	G int
	B int
}

// cmder + -> show helper
func main() {
	green := RGB{R: 130, G: 236, B: 99}
	PrintCustomln(green, "Hulu Hala day!")
	// fmt.Println(FormatBlueln("This is Blue"))
	// fmt.Println(FormatBlackln("This is Black"))
	// fmt.Println(FormatRedln("This is Red"))
	// fmt.Println(FormatGreenln("This is Green"))
	// fmt.Println(FormatYellowln("This is Yellow"))
	// fmt.Println(FormatMagentaln("This is Magenta"))
	// fmt.Println(FormatCyanln("This is Cyan"))
	// fmt.Println(FormatWhiteln("This is White"))
}
