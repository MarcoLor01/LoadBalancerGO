package wordCounter

import "fmt"

type Args struct {
	A, B string
}
type ResultCounter struct {
	LenA, LenB int
}

type Counter string

func (t *Counter) CountLettersReal(args Args, counter *ResultCounter) error {
	fmt.Print("I'm serving the request\n")
	counter.LenA = len(args.A)
	counter.LenB = len(args.B)
	return nil
}
