package main

import "github.com/yankeguo/weft"

// Modules register from init. Activate one with a blank import:
//
//	import (
//		"github.com/yankeguo/weft"
//		_ "github.com/custom/module"
//	)
func main() {
	weft.Main()
}
