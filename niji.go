package main

/*
Niji - 2018
* Does Work in Bash & ConEmu
! Does not Work in Powershell & Command Prompt

*/

// ESCAPECHAR Escape chars ver 1
var ESCAPECHAR = "\u001b["

// ESCAPECHAR2 Escape chars ver 2
var ESCAPECHAR2 = "\x1b["

// ESCAPECHAR3 Escape chars ver 3
var ESCAPECHAR3 = "\033["

var ESCAPECHARCLOSE = "\033[0m"

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
