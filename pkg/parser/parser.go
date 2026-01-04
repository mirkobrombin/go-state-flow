package parser

import (
	"fmt"
	"strings"
	"time"
)

type Config struct {
	InitialState string
	Transitions  map[string][]string
	Wildcards    []string
	Timeouts     map[string]TimeoutRule
}

type TimeoutRule struct {
	FromState string
	ToState   string
	Duration  time.Duration
}

func Parse(tag string) (*Config, error) {
	cfg := &Config{
		Transitions: make(map[string][]string),
		Timeouts:    make(map[string]TimeoutRule),
	}

	for part := range strings.SplitSeq(tag, ";") {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}

		if strings.HasPrefix(part, "initial:") {
			cfg.InitialState = strings.TrimPrefix(part, "initial:")
			continue
		}

		var timeoutDuration time.Duration
		if startBracket := strings.Index(part, "["); startBracket != -1 {
			endBracket := strings.Index(part, "]")
			if endBracket > startBracket {
				durStr := part[startBracket+1 : endBracket]
				var err error
				timeoutDuration, err = time.ParseDuration(durStr)
				if err != nil {
					return nil, fmt.Errorf("invalid timeout duration: %s", durStr)
				}
				part = strings.TrimSpace(part[:startBracket])
			}
		}

		before, after, ok := strings.Cut(part, "->")
		if !ok {
			return nil, fmt.Errorf("invalid transition syntax: %s", part)
		}

		src := strings.TrimSpace(before)
		dst := strings.TrimSpace(after)

		if src == "*" {
			cfg.Wildcards = append(cfg.Wildcards, dst)
		} else {
			cfg.Transitions[src] = append(cfg.Transitions[src], dst)
		}

		if timeoutDuration > 0 {
			cfg.Timeouts[src] = TimeoutRule{
				FromState: src,
				ToState:   dst,
				Duration:  timeoutDuration,
			}
		}
	}
	return cfg, nil
}
