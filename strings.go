package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "  string with spaces   \t\r\n"
	fmt.Println("|" + strings.TrimSpace(s) + "|")
}
