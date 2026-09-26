# Weftfile is the only config

A running Binary is configured only by a Weftfile. There is no JSON document underneath it, and no adapter step that compiles the DSL into another config. One format keeps each Module to a single settings surface.

## Considered Options

- **JSON as the real config, Weftfile as an adapter.** This is Caddy's split. Every Module would implement two config surfaces, and the DSL would not be the thing the process actually runs.

## Consequences

- A Module's settings have to be expressible in the Weftfile. There is no second config API.
- Changing the DSL later touches every Module's configuration.
