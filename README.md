# gocol
Super simple package for printing coloured text using readable tags.

## Get started
Firstly install the package with `go get github.com/DenizBaskan/gocol`. \
By default, text is written to `os.Stdout` however you can supply a different writer through `gocol.SetWriter` if needed.
### Example
```go
package main

import (
	"fmt"
	"os"

	"github.com/DenizBaskan/gocol"
)

func main() {
	// Optionally call gocol.SetWriter to write to os.Stderr instead
	gocol.SetWriter(os.Stderr)

	name := "Jack"
	// gocol.Colorln appends a reset and newline automatically
	gocol.Colorln("[bgblack][brightmagenta]Hello [bold]%s", name)

	gocol.Colorln("[red]Error")

	// Manually format colors in string and then call fmt.Printf appending a reset ansi code and a newline
	str := gocol.Sprintf("[green]I am green")
	fmt.Printf("%s\u001b[0m\n", str)
}
```
### Tags
```
{
		"[black]":   "\u001b[30m",
		"[red]":     "\u001b[31m",
		"[green]":   "\u001b[32m",
		"[yellow]":  "\u001b[33m",
		"[blue]":    "\u001b[34m",
		"[magenta]": "\u001b[35m",
		"[cyan]":    "\u001b[36m",
		"[white]":   "\u001b[37m",

		"[brightblack]":   "\u001b[90m",
		"[brightred]":     "\u001b[91m",
		"[brightgreen]":   "\u001b[92m",
		"[brightyellow]":  "\u001b[93m",
		"[brightblue]":    "\u001b[94m",
		"[brightmagenta]": "\u001b[95m",
		"[brightcyan]":    "\u001b[96m",
		"[brightwhite]":   "\u001b[97m",

		"[bgblack]":   "\u001b[40m",
		"[bgred]":     "\u001b[41m",
		"[bggreen]":   "\u001b[42m",
		"[bgyellow]":  "\u001b[43m",
		"[bgblue]":    "\u001b[44m",
		"[bgmagenta]": "\u001b[45m",
		"[bgcyan]":    "\u001b[46m",
		"[bgwhite]":   "\u001b[47m",

		"[bgbrightblack]":   "\u001b[100m",
		"[bgbrightred]":     "\u001b[101m",
		"[bgbrightgreen]":   "\u001b[102m",
		"[bgbrightyellow]":  "\u001b[103m",
		"[bgbrightblue]":    "\u001b[104m",
		"[bgbrightmagenta]": "\u001b[105m",
		"[bgbrightcyan]":    "\u001b[106m",
		"[bgbrightwhite]":   "\u001b[107m",

		"[reset]":     "\u001b[0m",
		"[bold]":      "\u001b[1m",
		"[dim]":       "\u001b[2m",
		"[underline]": "\u001b[4m",
		"[blink]":     "\u001b[5m",
		"[reversed]":  "\u001b[7m",
		"[hidden]":    "\u001b[8m",

		"[resetbold]":      "\u001b[21m",
		"[resetdim]":       "\u001b[22m",
		"[resetunderline]": "\u001b[24m",
		"[resetblink]":     "\u001b[25m",
		"[resetreversed]":  "\u001b[27m",
		"[resethidden]":    "\u001b[28m",
		"[resetbg]":        "\u001b[49m",
	}
```
