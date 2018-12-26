# Niji

Golang Coloring Cli Output.

## Installing

```bash
$ go get github.com/Jexulie/Niji
```

## How to Use
```golang
   // import package
   import niji "github.com/Jexulie/Niji"

   func main() {
    /* Standart colors
        Black,
        Red,
        Green,
        Yellow,
        Blue,
        Magenta,
        Cyan,
        White
   */

      // returns string
      str := niji.FormatGreen("Cow")
   
      // prints
      niji.PrintMagenta("Powerpuff Girls...")

   /* Custom Colors */

      // for HEX 
      blue := niji.HEX("#4D9FF7")

      // or as short form #49F

      niji.PrintCustomHEXln(blue, "Sky")

      // for RGB
      red := niji.RGB{R: 255, G: 0, B: 0}

      niji.PrintCustomRGBln(red, "Rawrr!")
   }


```

## Known Problems

* Program works in Bash, ConEmu & Git-Bash But Not in Powershell and Command Prompt

## Todo (Maybe)

- [ ] Support for Powershell and Command Prompt.
- [ ] Support for Background Coloring.
- [ ] Support for other stylings e.g. Underline, Italic, Bold, Strikethrough.

## [Licence MIT](https://github.com/Jexulie/Niji/blob/master/format.go)
