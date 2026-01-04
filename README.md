# Go State Flow

**Go State Flow** is a high-performance, declarative Finite State Machine (FSM) library for Go.

It solves the problem of "Spaghetti State Management" by allowing you to define valid states and transitions directly on your data structs using tags, rather than burying logic in `switch` statements or imperative code.

## Features

- **Declarative**: Define transitions in struct tags (e.g., `fsm:"draft->paid"`).
- **Hooks**: Automatic discovery of `OnEnter`, `OnExit`, and `CanEnter` (Guard) methods.
- **Zero Allocations**: Optimized runtime with method caching for high-performance hot paths.
- **Visualization**: Built-in export to Mermaid.js and Graphviz.

## Documentation

- **[Getting Started](docs/getting-started.md)**: Installation and your first FSM.
- **[Struct Tags](docs/struct-tags.md)**: Reference for the tag syntax.
- **[Hooks](docs/hooks.md)**: How to implement logic (`Can`, `OnEnter`, `OnExit`).
- **[Visualization](docs/visualization.md)**: Generating diagrams from your code.

## Quick Look

```go
type Order struct {
    // 1. Define the Machine
    Status string `fsm:"initial:draft; draft->paid; *->cancelled"`
}

// 2. Define Logic
func (o *Order) OnEnterPaid() {
    fmt.Println("Money received!")
}

// 3. Run
func main() {
    state, _ := stateflow.New(&Order{})
    state.Transition("paid")
}
```

## License

MIT License.
