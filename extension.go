package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	fmt.Println(filepath.Ext("a.jpg"))
	fmt.Println(filepath.Ext("/tmp/x/y.jpg/zg"))
	fmt.Println(filepath.Ext("/tmp/x/y.jpg/zg.another.one.here"))
}
