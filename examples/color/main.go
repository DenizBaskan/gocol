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
