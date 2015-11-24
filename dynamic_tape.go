package main

import (
	"fmt"
	"io"
	"unicode"
)

const size = 1024

type DynamicTape struct {
	cell []int
	pos  int
	out  io.ReadWriter
}

func NewDynamicTape(out io.ReadWriter) Storage {
	return &DynamicTape{
		cell: make([]int, size),
		pos:  0,
		out:  out,
	}
}

func (t *DynamicTape) check(pos int) {
	if pos >= len(t.cell) {
		t.cell = append(t.cell, make([]int, size)...)
	}
	if pos < 0 {
		t.cell = append(make([]int, size), t.cell...)
		t.pos += size
	}
}

func (t *DynamicTape) Move(n int) {
	t.pos += n
	t.check(t.pos)
}

func (t *DynamicTape) Add(n, off int) {
	x := t.pos + off
	t.check(x)
	t.cell[x] += n
}

func (t *DynamicTape) Print(off int) {
	x := t.pos + off
	t.check(x)
	if c := t.cell[x]; c > unicode.MaxASCII {
		fmt.Fprintf(t.out, "%d", c)
	} else {
		fmt.Fprintf(t.out, "%c", c)
	}
}

func (t *DynamicTape) Scan(off int) {
	x := t.pos + off
	t.check(x)
	fmt.Fscanf(t.out, "%c", &t.cell[x])
}

func (t *DynamicTape) IsZero() bool {
	return t.cell[t.pos] == 0
}

func (t *DynamicTape) Clear(off int) {
	x := t.pos + off
	t.check(x)
	t.cell[x] = 0
}

func (t *DynamicTape) Mult(dst, arg, off int) {
	x := t.pos + off
	t.check(x)
	v := t.cell[x]
	t.Move(dst)
	t.Add(v*arg, off)
	t.Move(-dst)
	//t.Clear() // inserted by optimization
}

func (t *DynamicTape) Search(n int) {
	for !t.IsZero() {
		t.Move(n)
	}
}

func (t *DynamicTape) Dump() ([]int, int) {
	return t.cell, t.pos
}
