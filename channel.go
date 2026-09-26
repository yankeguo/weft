package weft

import "context"

// Channel is where a person's messages enter a session and where replies leave.
// A channel is a module: RegisterChannel links it into the binary at build time.
type Channel interface {
	// Run connects the channel and blocks until ctx is canceled or the channel fails.
	// handle runs one turn. Core queues a later call with the same session key
	// until the earlier call returns. The channel does not do that queueing.
	Run(ctx context.Context, handle TurnHandler) error
}

// TurnHandler runs one turn for an incoming message.
type TurnHandler func(ctx context.Context, in Incoming) (Outgoing, error)

// Incoming is one user message, addressed to a session.
// SessionKey is opaque to Core. The channel chooses it.
type Incoming struct {
	SessionKey string
	Text       string
}

// Outgoing is the reply Core returns to the channel for one turn.
type Outgoing struct {
	Text string
}

// RegisterChannel registers a channel module under name.
// Call it from init. It panics if name is empty, newChannel is nil,
// the name is already registered, or Main has already started.
func RegisterChannel(name string, newChannel func() Channel) {
	defaultRegistry.registerChannel(name, newChannel)
}
