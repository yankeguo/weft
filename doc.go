// Package weft is the agent core.
//
// A binary is this package plus the modules linked into it. Channels and
// direct tools register themselves from init:
//
//	func init() {
//	    weft.RegisterChannel("discord", func() weft.Channel { return Discord{} })
//	    weft.RegisterTool("query_orders", func() weft.Tool { return QueryOrders{} })
//	}
//
// A custom build activates those modules with a blank import in cmd/weft/main.go:
//
//	import (
//	    "github.com/yankeguo/weft"
//	    _ "github.com/custom/module"
//	)
//
// Model protocols and MCP are part of Core. They are not registered this way.
package weft
