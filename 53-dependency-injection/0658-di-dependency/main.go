package main

import (
	"fmt"

	"go.uber.org/dig"
)

// Repo is a leaf dependency.
type Repo struct{}

func (r *Repo) Data() string { return "data" }

// Service depends on *Repo and delegates run() to it.
type Service struct{ repo *Repo }

func (s *Service) Run() string { return s.repo.Data() }

// Constructors are the providers dig wires together.
func NewRepo() *Repo { return &Repo{} }

func NewService(r *Repo) *Service { return &Service{repo: r} }

func main() {
	c := dig.New()
	// Register both providers; dig resolves NewService's *Repo
	// argument from NewRepo automatically.
	if err := c.Provide(NewRepo); err != nil {
		panic(err)
	}
	if err := c.Provide(NewService); err != nil {
		panic(err)
	}
	// Resolve Service (which pulls in Repo) and call run().
	if err := c.Invoke(func(s *Service) {
		fmt.Println(s.Run())
	}); err != nil {
		panic(err)
	}
}
