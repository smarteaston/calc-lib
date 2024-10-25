package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/smarteaston/calc-lib/calc"
)

type Calculator interface {
	Calculate(a, b int) int
}

func main() {
	handler := NewHandler(os.Stdout, calc.Addition{})
	err := handler.handle(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
}

type Handler struct {
	output     io.Writer
	calculator Calculator
}

func NewHandler(out io.Writer, calculator Calculator) *Handler {
	return &Handler{
		output:     out,
		calculator: calculator,
	}
}

func (this *Handler) handle(args []string) error {
	if len(args) != 2 {
		return errWrongArgCount
	}
	a, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("%w: %s", errInvalidArg, err)
	}
	b, err := strconv.Atoi(args[1])
	if err != nil {
		return fmt.Errorf("%w: %s", errInvalidArg, err)
	}
	_, err = fmt.Fprint(this.output, this.calculator.Calculate(a, b))
	return err
}

var (
	errWrongArgCount = errors.New("wrong number of arguments")
	errInvalidArg    = errors.New("invalid argument")
)
