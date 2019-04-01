// Copyright 2019 The color Authors. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package color

const (
	allocMinSize = 0
	allocMaxSize = 2
)

type pool struct {
	c chan *Color
}

func newColorPool(size int) (p *pool) {
	return &pool{
		c: make(chan *Color, size),
	}
}

func (p *pool) Get() (c *Color) {
	select {
	case c = <-p.c:
		// reuse existing *Color
	default:
		c = &Color{params: make([]Attribute, allocMinSize, allocMaxSize)}
	}

	return
}

func (p *pool) Put(c *Color) {
	c.Reset()

	select {
	case p.c <- c:
	default:
		// pool is full. discard the c.
	}
}
