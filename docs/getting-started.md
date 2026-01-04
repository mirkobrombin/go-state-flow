# Getting Started with Go State Flow

Go State Flow allows you to add state machine capabilities to your existing Go structs with minimal friction.

## Installation

```bash
go get github.com/mirkobrombin/go-state-flow
```

## Basic Concept

The library works by reading a special `fsm` struct tag on your data model. It uses this tag to build a directed graph of valid state transitions.

When you call `machine.Transition(newState)`, the library:
1. Checks if the transition is valid.
2. Runs a **Guard** (`Can<State>`) to see if the business logic allows it.
3. Runs an **Exit** hook (`OnExit<CurrentState>`) for cleanup.
4. Updates the state field.
5. Runs an **Enter** hook (`OnEnter<NewState>`) for side effects.

## Quick Example

Here is a complete (damn easy), runnable example of a simple "Switch" state machine.
    stateflow "github.com/mirkobrombin/go-state-flow/pkg/machine"
```go
package main

import (
    "fmt"
    stateflow "github.com/mirkobrombin/go-state-flow/pkg/core"
)

type Switch struct {
    // Define the state machine here.
    // We start at 'off'.
    // We can go from 'off' to 'on', and 'on' to 'off'.
    State string `fsm:"initial:off; off->on; on->off"`
}

// Hook: Called when we enter the 'on' state.
func (s *Switch) OnEnterOn() {
    fmt.Println("Lights are ON!")
}

// Hook: Called when we enter the 'off' state.
func (s *Switch) OnEnterOff() {
    fmt.Println("Lights are OFF.")
}

func main() {
    sw := &Switch{}
    fsm, _ := stateflow.New(sw)

    fsm.Transition("on")  // Output: Lights are ON!
    fsm.Transition("off") // Output: Lights are OFF.
}
```

Now imagine the same code without this library:

```go
package main

import "fmt"

type Switch struct {
    State string
}

func (s *Switch) OnEnterOn() {
    fmt.Println("Lights are ON!")
}

func (s *Switch) OnEnterOff() {
    fmt.Println("Lights are OFF.")
}

func main() {
    sw := &Switch{}
    if sw.State != "on" {
        sw.State = "on"
        sw.OnEnterOn()
    }
    if sw.State != "off" {
        sw.State = "off"
        sw.OnEnterOff()
    }
}
```

## Next Steps

- Learn about [Struct Tags](struct-tags.md) syntax.
- Understand [Hooks](hooks.md) for business logic.
- See how to [Visualize](visualization.md) your FSM.
