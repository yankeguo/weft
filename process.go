package weft

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

const usage = `Usage: weft

Run this binary. Channels and direct tools are modules linked at build time.
Add a module by blank-importing it from cmd/weft/main.go. Its init calls
weft.RegisterChannel or weft.RegisterTool.
`

// Main runs the weft process.
// Module init functions have already run, including blank imports from the
// command in cmd/weft.
func Main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	if err := run(ctx, os.Args[1:], os.Stderr); err != nil {
		fmt.Fprintf(os.Stderr, "weft: %v\n", err)
		os.Exit(1)
	}
}

func run(ctx context.Context, args []string, stderr io.Writer) error {
	if len(args) == 1 && (args[0] == "-h" || args[0] == "--help" || args[0] == "help") {
		fmt.Fprint(stderr, usage)
		return nil
	}
	if len(args) != 0 {
		return fmt.Errorf("unexpected arguments: %s", strings.Join(args, " "))
	}

	defaultRegistry.markStarted()
	channels, tools := defaultRegistry.counts()
	fmt.Fprintf(stderr, "weft: ready channels=%d tools=%d\n", channels, tools)
	<-ctx.Done()
	return nil
}
