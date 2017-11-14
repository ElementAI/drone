package main

import (
	"io"

	"github.com/cncd/pipeline/pipeline/backend"
)

type Borgy struct {
	c *borgy.Client
}

// Setup the pipeline environment.
func (b *Borgy) Setup(c *backend.Config) error {
	return nil
}

// Start the pipeline step.
func (b *Borgy) Exec(s *backend.Step) error {
	return nil
}

// Kill the pipeline step.
func (b *Borgy) Kill(s *backend.Step) error {
	return nil
}

// Wait for the pipeline step to complete and returns the completion results.
func (b *Borgy) Wait(s *backend.Step) (*backend.State, error) {
	return nil, nil
}

// Tail the pipeline step logs.
func (b *Borgy) Tail(s *backend.Step) (io.ReadCloser, error) {
	return nil, nil
}

// Destroy the pipeline environment.
func (b *Borgy) Destroy(c *backend.Config) error {
	return nil
}
