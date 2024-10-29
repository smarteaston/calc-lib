package main

import (
	"bytes"
	"errors"
	"strings"
	"testing"

	"github.com/smarteaston/calc-lib/calc"
)

func assertErr(t *testing.T, actual error, targets ...error) {
	for _, target := range targets {
		if !errors.Is(actual, target) {
			t.Errorf("wanted %v, got %v", target, actual)
		}
	}
}

func TestHandler_TwoArgsRequired(t *testing.T) {
	handler := NewHandler(nil, nil)
	err := handler.handle(nil)
	assertErr(t, err, errWrongArgCount)
}

func TestHandler_FirstArgInvalid(t *testing.T) {
	handler := NewHandler(nil, nil)
	err := handler.handle([]string{"hi", "42"})
	assertErr(t, err, errInvalidArg)
}

func TestHandler_SecondArgInvalid(t *testing.T) {
	handler := NewHandler(nil, nil)
	err := handler.handle([]string{"42", "hi"})
	if !errors.Is(err, errInvalidArg) {
		t.Errorf("want: %v, got: %v", errInvalidArg, err)
	}
}

func TestHandler_InvalidFP(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected a nil pointer")
		}
	}()
	handler := NewHandler(nil, calc.Addition{})
	handler.handle([]string{"1", "2"})
}

func TestHandler_InvalidCalculator(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected a nil pointer")
		}
	}()
	handler := NewHandler(&strings.Builder{}, nil)
	handler.handle([]string{"1", "2"})
}

func TestHandler_ResultsWrittenToOutput(t *testing.T) {
	stdout := bytes.Buffer{}
	handler := NewHandler(&stdout, calc.Addition{})
	err := handler.handle([]string{"1", "2"})
	assertErr(t, err, nil)
	if stdout.String() != "3" {
		t.Errorf("want: %s, got: %s", "3", stdout.String())
	}
}

type BadWriter struct {
	err error
}

func (this BadWriter) Write(p []byte) (n int, err error) {
	return 0, this.err
}

func TestHandler_OutputtingError(t *testing.T) {
	nico := errors.New("nico and the niners")
	badWriter := &BadWriter{err: nico}
	handler := NewHandler(badWriter, calc.Addition{})
	err := handler.handle([]string{"5", "5"})
	assertErr(t, err, nico, errOutputProblem)
}
