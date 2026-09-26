# Weft

Weft is an agent host. Each binary runs a conversation loop, and the channels and direct tools it can use are fixed when that binary is built. A Weftfile is the only way to configure a running binary.

## Language

### Composition

**Binary**:
A compiled program made of Core plus a fixed set of Modules. A Weftfile can name only Modules that are in that Binary.
_Avoid_: Distribution, plugin host

**Core**:
The part present in every Binary: the Loop, the Weftfile, the built-in model protocols, and MCP.
_Avoid_: Runtime, framework, kernel

**Module**:
A named capability compiled into a Binary. Channels and direct tools are Modules.
_Avoid_: Plugin, extension, addon

**Weftfile**:
The only configuration for a running Binary. It selects Modules and records their settings.
_Avoid_: JSON config, manifest

### Conversation

**Channel**:
Where a person's messages enter a Session and where replies leave.
_Avoid_: Connector, integration, ingress

**Session**:
One conversation the Loop runs. A message that arrives during a Turn waits on that Session until the Turn finishes, then becomes the next Turn.
_Avoid_: Thread, chat, context

**Turn**:
The span from one user message being taken from the Session until a reply is ready to send back.
_Avoid_: Request, run, completion

**Loop**:
The Core behavior that executes a Turn against a Model and a set of Tools.
_Avoid_: Agent, orchestrator

### Capabilities

**Tool**:
A capability the Loop may call during a Turn. Direct tools and tools from MCP servers are the same kind of thing to the Loop.
_Avoid_: Function, skill, command

**Direct tool**:
A Tool whose implementation is a Module in the Binary.
_Avoid_: Native tool, builtin tool

**MCP server**:
A separate process that offers Tools over MCP. Core connects to it, and those Tools join the same set as direct tools.
_Avoid_: Plugin process, sidecar

**Model**:
An endpoint the Loop talks to, using one built-in Protocol.
_Avoid_: LLM, provider, backend

**Protocol**:
One of the three wire formats Core speaks: OpenAI Chat Completions, OpenAI Responses, and Anthropic.
_Avoid_: Provider, vendor, driver
