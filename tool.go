package weft

import (
	"context"
	"encoding/json"
)

// Tool is a capability the loop may call during a turn.
// RegisterTool links a direct tool into the binary. Tools that come from
// an MCP server are not registered here.
type Tool interface {
	// Description is the text the model sees for this tool.
	Description() string
	// Parameters is a JSON Schema object for the arguments.
	// A nil schema means the tool takes no arguments.
	Parameters() json.RawMessage
	// Execute runs the tool. The result is the text the model sees.
	Execute(ctx context.Context, args json.RawMessage) (string, error)
}

// RegisterTool registers a direct tool module under name.
// Call it from init. It panics if name is empty, newTool is nil,
// the name is already registered, or Main has already started.
func RegisterTool(name string, newTool func() Tool) {
	defaultRegistry.registerTool(name, newTool)
}
