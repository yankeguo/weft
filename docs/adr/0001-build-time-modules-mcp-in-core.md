# Build-time modules, with MCP kept in Core

Channels and direct tools are Modules linked into a Binary at build time. Core still speaks MCP, so an MCP server's tools join the same set the Loop calls. A separate process and protocol is the wrong default for a tool Weft itself should carry, and existing MCP servers should still work without being rewritten as Modules first.

## Considered Options

- **MCP only.** Every tool is a separate process. This is the setup Weft exists to move past for tools it can compile in.
- **Compile-time modules only.** Drop MCP from Core. That forces every existing server to be rewritten before it can be used.
- **Dynamic loading.** Drop a plugin into a directory at runtime. That reintroduces a runtime composition step and a plugin ABI.

## Consequences

- A Weftfile can name only Modules that were linked into that Binary. A missing name fails at load.
- A custom deployment is a custom Binary: Core plus the Channels and direct tools that build imports.
- Direct tools and MCP servers both contribute Tools. How to resolve the same name from both is not decided yet.
