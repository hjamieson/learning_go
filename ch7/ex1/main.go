package main

import (
	"fmt"
	"time"
)

type Counter struct {
	total int
	lastUpdated time.Time
}

func (c* Counter) Increment() {
	c.total++
	c.lastUpdated = time.Now()
}

func (c Counter) String() string {
	return fmt.Sprintf("Total: %d, Last Updated: %v", c.total, c.lastUpdated)
}

type Incrementer interface {
	Increment()
}

func proc2() {

	var myStringer fmt.Stringer
	var myIncrementer Incrementer
	pointerCounter := &Counter{}
	valueCounter := Counter{}

	myStringer = pointerCounter
	myStringer = valueCounter
	myStringer = pointerCounter
	myIncrementer = valueCounter
}