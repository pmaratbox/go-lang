package main

import (
	"fmt"

	"go.uber.org/dig"
)

// Repo is the service we register as a singleton.
type Repo struct{}

func (r *Repo) Data() string { return "data" }

// NewRepo is the constructor dig calls to build a *Repo.
func NewRepo() *Repo { return &Repo{} }

func main() {
	c := dig.New()
	c.Provide(NewRepo)

	// Resolve the same dependency twice. dig caches the value produced by a
	// constructor, so both resolutions hand back the exact same *Repo pointer.
	var p1, p2 *Repo
	c.Invoke(func(r *Repo) { p1 = r })
	c.Invoke(func(r *Repo) { p2 = r })

	// Singleton lifetime => identical instance => identity is true.
	fmt.Println(p1 == p2)
}
