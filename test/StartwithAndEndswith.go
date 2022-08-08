package main

import (
	"strings"
)

func main() {
	startsWith := strings.HasPrefix("prefix", "pre")
	endsWith := strings.HasSuffix("suffix", "fix")
	println(startsWith)
	println(endsWith)
}
