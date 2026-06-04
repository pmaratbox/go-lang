package main

import "fmt"

type subject interface {
	request() string
}

type realSubject struct{}

func (realSubject) request() string { return "loaded" }

type proxy struct{ real *realSubject }

func (p *proxy) request() string {
	if p.real == nil {
		p.real = &realSubject{}
	}
	return p.real.request()
}

func main() {
	var s subject = &proxy{}
	fmt.Println(s.request())
}
