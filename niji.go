package niji

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

// RGBControl c
func (r RGB) RGBControl() {
	if r.R < 0 || r.G < 0 || r.B < 0 {
		panic("Color Values Can Not Be Negative")
	}
	if r.R > 255 || r.G > 255 || r.B > 255 {
		panic("Color Values Can Not Be Higher Than 255")
	}
}

// HEX struct
type HEX string

// HEXControl c
func HEXControl(s []string) {
	if len(s) < 3 {
		panic("Something is Wrong with your Hex!")
	}
}

// ToRGB h -> r
func (h HEX) ToRGB() RGB {
	var x RGB
	if len(h) == 7 || len(h) == 4 {
		hex := `(?i)^#([a-f\d]{2})([a-f\d]{2})([a-f\d]{2})`
		comp := regexp.MustCompile(hex)
		r := comp.FindStringSubmatch(string(h))
		HEXControl(r)
		tr, _ := strconv.ParseInt(r[1], 16, 0)
		tg, _ := strconv.ParseInt(r[2], 16, 0)
		tb, _ := strconv.ParseInt(r[3], 16, 0)
		x.R = int(tr)
		x.G = int(tg)
		x.B = int(tb)
	}
	return x
}
