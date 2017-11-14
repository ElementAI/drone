package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"strings"

	"github.com/ElementAI/go-borgy/borgy"
	"github.com/cncd/pipeline/pipeline/backend"
)

func NewBorgy(user string) *Borgy {
	return &Borgy{
		c: &borgy.Client{
			User: user,
		},
	}
}

type Borgy struct {
	c *borgy.Client
}

// Setup the pipeline environment.
func (b *Borgy) Setup(c *backend.Config) error {
	js, _ := json.MarshalIndent(c, "", "  ")
	fmt.Printf("Setup: %s\n", js)
	return nil
}

// Start the pipeline step.
func (b *Borgy) Exec(s *backend.Step) error {
	js, _ := json.MarshalIndent(s, "", "  ")
	fmt.Printf("Exec: %s\n", js)

	env := []string{}
	for k, v := range s.Environment {
		env = append(env, fmt.Sprintf("%s=%s", k, v))
	}

	args := borgy.Template{
		Command:      s.Command,
		Environment:  env,
		Image:        s.Image,
		Name:         s.Name,
		RequiredCPUs: 1,
		Volumes:      s.Volumes,
	}

	job, err := b.c.Start(args)
	if err != nil {
		return err
	}

	fmt.Println(job)
	return nil
}

// Kill the pipeline step.
func (b *Borgy) Kill(s *backend.Step) error {
	js, _ := json.MarshalIndent(s, "", "  ")
	fmt.Printf("Kill: %s\n", js)
	return nil
}

// Wait for the pipeline step to complete and returns the completion results.
func (b *Borgy) Wait(s *backend.Step) (*backend.State, error) {
	js, _ := json.MarshalIndent(s, "", "  ")
	fmt.Printf("Wait: %s\n", js)
	return &backend.State{ExitCode: 1, Exited: true}, nil
}

// Tail the pipeline step logs.
func (b *Borgy) Tail(s *backend.Step) (io.ReadCloser, error) {
	js, _ := json.MarshalIndent(s, "", "  ")
	fmt.Printf("Tail: %s\n", js)
	return ioutil.NopCloser(strings.NewReader("tail\nhello\nworld\n")), nil
}

// Destroy the pipeline environment.
func (b *Borgy) Destroy(c *backend.Config) error {
	js, _ := json.MarshalIndent(c, "", "  ")
	fmt.Printf("Destroy: %s\n", js)
	return nil
}
