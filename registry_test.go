package weft

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"testing"
)

func resetRegistry() {
	defaultRegistry.mu.Lock()
	defer defaultRegistry.mu.Unlock()
	defaultRegistry.started = false
	defaultRegistry.channels = map[string]func() Channel{}
	defaultRegistry.tools = map[string]func() Tool{}
}

func mustPanic(t *testing.T, substr string, fn func()) {
	t.Helper()
	defer func() {
		r := recover()
		if r == nil {
			t.Fatal("expected panic")
		}
		if !strings.Contains(fmt.Sprint(r), substr) {
			t.Fatalf("panic %v, want substring %q", r, substr)
		}
	}()
	fn()
}

type nopChannel struct{}

func (nopChannel) Run(context.Context, TurnHandler) error { return nil }

type nopTool struct{}

func (nopTool) Description() string         { return "" }
func (nopTool) Parameters() json.RawMessage { return nil }
func (nopTool) Execute(context.Context, json.RawMessage) (string, error) {
	return "", nil
}

func TestRegisterChannelAndTool(t *testing.T) {
	t.Cleanup(resetRegistry)

	RegisterChannel("discord", func() Channel { return nopChannel{} })
	RegisterTool("discord", func() Tool { return nopTool{} })

	channels, tools := defaultRegistry.counts()
	if channels != 1 || tools != 1 {
		t.Fatalf("counts = %d channels, %d tools", channels, tools)
	}
}

func TestRegisterRejectsDuplicatesAndEmptyValues(t *testing.T) {
	t.Cleanup(resetRegistry)

	RegisterChannel("discord", func() Channel { return nopChannel{} })
	mustPanic(t, "already registered", func() {
		RegisterChannel("discord", func() Channel { return nopChannel{} })
	})
	mustPanic(t, "name is empty", func() {
		RegisterChannel("", func() Channel { return nopChannel{} })
	})
	mustPanic(t, "constructor is nil", func() {
		RegisterChannel("cli", nil)
	})

	RegisterTool("search", func() Tool { return nopTool{} })
	mustPanic(t, "already registered", func() {
		RegisterTool("search", func() Tool { return nopTool{} })
	})
	mustPanic(t, "name is empty", func() {
		RegisterTool("", func() Tool { return nopTool{} })
	})
	mustPanic(t, "constructor is nil", func() {
		RegisterTool("shell", nil)
	})
}

func TestRegisterAfterStartPanics(t *testing.T) {
	t.Cleanup(resetRegistry)
	defaultRegistry.markStarted()

	mustPanic(t, "after start", func() {
		RegisterChannel("discord", func() Channel { return nopChannel{} })
	})
	mustPanic(t, "after start", func() {
		RegisterTool("search", func() Tool { return nopTool{} })
	})
}
