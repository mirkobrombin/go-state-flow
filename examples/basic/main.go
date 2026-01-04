package main

import (
	"fmt"

	"github.com/mirkobrombin/go-state-flow/pkg/machine"
)

type Order struct {
	ID     string
	Status string `fsm:"initial:draft; draft->pending; pending->paid; paid->shipped; *->cancelled"`
}

func main() {
	order := &Order{ID: "ORD-123"}

	m, err := machine.New(order)
	if err != nil {
		panic(err)
	}

	fmt.Printf("State: %s\n", order.Status)

	if err := m.Transition("pending"); err != nil {
		panic(err)
	}
	fmt.Printf("State: %s\n", order.Status)

	if err := m.Transition("shipped"); err == nil {
		fmt.Println("Error: Expected transition failure, but got success")
	} else {
		fmt.Printf("Blocked: %v\n", err)
	}
}
