# Struct Tags Syntax

The core of Go State Flow is the declarative `fsm` struct tag. This tag defines the topology of your state machine.

## Format

The tag is a semicolon-separated list of instructions.

```go
`fsm:"instruction1; instruction2; ..."`
```

## Instructions

### 1. Initial State
Sets the default state of the struct if the field is empty when `stateflow.New()` is called.

**Syntax:** `initial:<state_name>`

```go
`fsm:"initial:draft"`
```

### 2. Transition
Defines a valid path from one state to another.

**Syntax:** `<source>-><destination>`

```go
`fsm:"draft->published"`
```

### 3. Wildcard Transition
Defines a transition that is allowed from ANY state. This is useful for cancellation flows or error states.

**Syntax:** `*-><destination>`

```go
`fsm:"*->cancelled"`
```

## Example

```go
type Document struct {
    Status string `fsm:"initial:draft; draft->review; review->published; *->archived"`
}
```

This defines:
1.  Start at `draft`.
2.  `draft` can go to `review`.
3.  `review` can go to `published`.
4.  ANY state (`draft`, `review`, `published`) can go to `archived`.
