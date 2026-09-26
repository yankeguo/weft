package weft

import (
	"bytes"
	"context"
	"io"
	"strings"
	"testing"
)

func TestRunReady(t *testing.T) {
	t.Cleanup(resetRegistry)
	RegisterChannel("cli", func() Channel { return nopChannel{} })

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	var stderr bytes.Buffer
	if err := run(ctx, nil, &stderr); err != nil {
		t.Fatal(err)
	}
	got := stderr.String()
	if !strings.Contains(got, "channels=1") || !strings.Contains(got, "tools=0") {
		t.Fatalf("stderr = %q", got)
	}
}

func TestRunHelp(t *testing.T) {
	var stderr bytes.Buffer
	if err := run(context.Background(), []string{"--help"}, &stderr); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(stderr.String(), "RegisterChannel") {
		t.Fatalf("stderr = %q", stderr.String())
	}
}

func TestRunUnexpectedArgs(t *testing.T) {
	err := run(context.Background(), []string{"nope"}, io.Discard)
	if err == nil || !strings.Contains(err.Error(), "unexpected arguments") {
		t.Fatalf("err = %v", err)
	}
}
