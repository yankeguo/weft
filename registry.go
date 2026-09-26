package weft

import "sync"

// registry is the set of modules linked into this process.
// init functions register into it, and Main marks it started.
type registry struct {
	mu       sync.Mutex
	started  bool
	channels map[string]func() Channel
	tools    map[string]func() Tool
}

var defaultRegistry = &registry{
	channels: map[string]func() Channel{},
	tools:    map[string]func() Tool{},
}

func (r *registry) registerChannel(name string, newChannel func() Channel) {
	if name == "" {
		panic("weft: RegisterChannel name is empty")
	}
	if newChannel == nil {
		panic("weft: RegisterChannel constructor is nil")
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.started {
		panic("weft: RegisterChannel called after start")
	}
	if _, ok := r.channels[name]; ok {
		panic("weft: channel " + name + " is already registered")
	}
	r.channels[name] = newChannel
}

func (r *registry) registerTool(name string, newTool func() Tool) {
	if name == "" {
		panic("weft: RegisterTool name is empty")
	}
	if newTool == nil {
		panic("weft: RegisterTool constructor is nil")
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.started {
		panic("weft: RegisterTool called after start")
	}
	if _, ok := r.tools[name]; ok {
		panic("weft: tool " + name + " is already registered")
	}
	r.tools[name] = newTool
}

func (r *registry) markStarted() {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.started = true
}

func (r *registry) counts() (channels int, tools int) {
	r.mu.Lock()
	defer r.mu.Unlock()
	return len(r.channels), len(r.tools)
}
