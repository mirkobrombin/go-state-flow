package main

import (
	"fmt"
	"time"

	"github.com/mirkobrombin/go-state-flow/pkg/machine"
)

type Payment struct {
	ID     string
	Status string `fsm:"initial:pending; pending->paid; pending->expired [500ms]"`
}

func main() {
	p := &Payment{ID: "TX-999"}
	m, _ := machine.New(p)

	m.Subscribe(func(e machine.Event) {
		if e.Type == machine.EnterState {
			fmt.Printf("[Event] Entered state: %s (Triggered by %s)\n", e.To, e.From)
		}
	})

	fmt.Printf("Current: %s\n", m.CurrentState())

	fmt.Println("Waiting for timeout...")
	time.Sleep(600 * time.Millisecond)

	if err := m.CheckTimeouts(); err != nil {
		fmt.Println("Warning:", err)
	}

	fmt.Printf("Current: %s\n", m.CurrentState())

	fmt.Println("\n--- History ---")
	for _, rec := range m.History() {
		fmt.Printf("%s -> %s (at %s) [Trigger: %s]\n",
			rec.From, rec.To, rec.Timestamp.Format(time.TimeOnly), rec.Trigger)
	}
}
