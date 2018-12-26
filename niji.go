package main

import (
	"regexp"
	"strconv"
)

/*
* Niji - 2018
* Does Work in Bash & ConEmu & Git-Bash
! Does not Work in Powershell & Command Prompt
*/

// ESCAPECHAR Escape chars ver 1
var ESCAPECHAR = "\u001b["

// ESCAPECHARCLOSE Escape char close 1
var ESCAPECHARCLOSE = "\u001b[0m"

// ESCAPECHAR2 Escape chars ver 2
var ESCAPECHAR2 = "\x1b["

// ESCAPECHARCLOSE2 Escape char close 2
var ESCAPECHARCLOSE2 = "\x1b[0m"

// ESCAPECHAR3 Escape chars ver 3
var ESCAPECHAR3 = "\033["

// ESCAPECHARCLOSE3 Escape char close 3
var ESCAPECHARCLOSE3 = "\033[0m"

// Niji type
type Niji int

// Niji Enum
const (
	Black Niji = iota + 30
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

// HEX struct
type HEX string

// ToRGB h -> r
func (h HEX) ToRGB() RGB {
	var x RGB
	if len(h) == 7 || len(h) == 4 {
		hex := `(?i)^#([a-f\d]{2})([a-f\d]{2})([a-f\d]{2})`
		comp := regexp.MustCompile(hex)
		r := comp.FindStringSubmatch(string(h))
		tr, _ := strconv.ParseInt(r[1], 16, 0)
		tg, _ := strconv.ParseInt(r[2], 16, 0)
		tb, _ := strconv.ParseInt(r[3], 16, 0)
		x.R = int(tr)
		x.G = int(tg)
		x.B = int(tb)
	}
	return x
}

// warning type hexToRGB func(h HEX) RGB

// cmder + -> show helper
func main() {
	var a HEX = "#F78393"
	b := HEX("#85F4B1")
	c := HEX("#937CC4")
	d := RGB{R: 225, G: 225, B: 48}
	e := RGB{R: 41, G: 61, B: 135}
	f := RGB{R: 174, G: 171, B: 167}

	PrintCustomHEXln(a, "Sample Hex !")
	PrintCustomHEXln(b, "Sample Hex !")
	PrintCustomHEXln(c, "Sample Hex !")
	PrintCustomRGBln(d, "Sample RGB !")
	PrintCustomRGBln(e, "Sample RGB !")
	PrintCustomRGBln(f, "Sample RGB !")

}
