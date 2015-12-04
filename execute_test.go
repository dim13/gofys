package main

import (
	"os"
	"testing"
)

func exec(prog string) {
	t := NewStaticTape(os.Stdout)
	p := ParseString(prog)
	Execute(Optimize(p), t)
}

// Hello World!
const helloWorld = `++++++++++[>+++++++>++++++++++>+++>+<<<<-]
>++.>+.+++++++..+++.>++.<<+++++++++++++++.>.+++.------.--------.>+.>.`

func ExampleHelloWorld() {
	exec(helloWorld)
	// Output: Hello World!
}

// Prints 202
const numeric = `>+>+>+>+>++<[>[<+++>-]<<]>.`

// Numeric output
func ExampleNumeric() {
	exec(numeric)
	// Output: 202
}

// Goes to cell 30000 and reports from there with a #. (Verifies that the
// array is big enough.)
const faraway = `++++[>++++++<-]>[>+++++>+++++++<<-]>>++++<
[[>[[>>+<<-]<]>>>-]>-[>+>+<<-]>]+++++[>+++++++<<++>-]>.<<.`

func ExampleFarAway() {
	exec(faraway)
	// Output: #
}

type devNull struct{}

func (devNull) Write(p []byte) (int, error) { return len(p), nil }
func (devNull) Read(p []byte) (int, error)  { return len(p), nil }

func bench(b *testing.B, fname string, optimize bool) {
	p, err := ParseFile(fname)
	if err != nil {
		b.Fatal(err)
	}
	if optimize {
		p = Optimize(p)
	}
	for i := 0; i < b.N; i++ {
		t := NewStaticTape(devNull{})
		Execute(p, t)
	}
}

func BenchmarkHanoi(b *testing.B)         { bench(b, "samples/hanoi.b", true) }
func BenchmarkHanoiRaw(b *testing.B)      { bench(b, "samples/hanoi.b", false) }
func BenchmarkMandelbrot(b *testing.B)    { bench(b, "samples/mandelbrot.b", true) }
func BenchmarkMandelbrotRaw(b *testing.B) { bench(b, "samples/mandelbrot.b", false) }
func BenchmarkLong(b *testing.B)          { bench(b, "samples/long.b", true) }
func BenchmarkLongRaw(b *testing.B)       { bench(b, "samples/long.b", false) }
