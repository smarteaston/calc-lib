package main

import (
	"bytes"
	"errors"
	"strings"
	"testing"

	"github.com/smarteaston/calc-lib/calc"
)

func TestHandler_TwoArgsRequired(t *testing.T) {
	handler := NewHandler(nil, nil)
	err := handler.handle(nil)
	if !errors.Is(err, errWrongArgCount) {
		t.Errorf("want: %v, got: %v", errWrongArgCount, err)
	}
}

func TestHandler_FirstArgInvalid(t *testing.T) {
	handler := NewHandler(nil, nil)
	err := handler.handle([]string{"hi", "42"})
	if !errors.Is(err, errInvalidArg) {
		t.Errorf("want: %v, got: %v", errInvalidArg, err)
	}
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
	if err != nil {
		t.Errorf("wasn't expecting an error")
	}
	if stdout.String() != "3" {
		t.Errorf("want: %s, got: %s", "3", stdout.String())
	}
}
