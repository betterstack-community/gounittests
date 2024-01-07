package main

import (
	"errors"
)

var (
	ErrCounterExhausted = errors.New("counter exhausted")
	ErrNegativeCounter  = errors.New("counter cannot be negative")
)

type Counter int64

func NewCounter(i int64) *Counter {
	c := Counter(i)
	return &c
}

func (c *Counter) Value() int64 {
	if *c == 0 {
		return 0
	}

	return int64(*c)
}

func (c *Counter) Take() error { return c.TakeN(1) }

func (c *Counter) TakeN(i int64) error {
	if *c <= 0 {
		return ErrCounterExhausted
	}

	if remaining := *c - Counter(i); remaining < 0 {
		return ErrNegativeCounter
	} else {
		*c = remaining
	}

	return nil
}

type Entitlement struct {
	Teams struct {
		Enabled bool     `json:"enabled,omitempty"`
		Size    *Counter `json:"size,omitempty"`
	} `json:"teams,omitempty"`
}