package machine

import (
	"time"
)

type EventType int

const (
	BeforeTransition EventType = iota
	AfterTransition
	EnterState
	ExitState
)

type TransitionRecord struct {
	From      string
	To        string
	Timestamp time.Time
	Trigger   string
	Metadata  map[string]any
}

type Listener func(e Event)

type Event struct {
	Type      EventType
	From      string
	To        string
	Timestamp time.Time
	Machine   *Machine
}
