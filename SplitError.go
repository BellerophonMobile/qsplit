package qsplit

import (
	"fmt"
)

const (
	InvalidEscapedCharacter = iota
	UnterminatedQuote
)

type SplitError struct {
	Type int
	Data string
}

func (x SplitError) Error() string {
	switch x.Type {
	case InvalidEscapedCharacter:
		return fmt.Sprintf("Invalid escaped character '%v'", x.Data)
	case UnterminatedQuote:
		return fmt.Sprintf("Quote not terminated")
	}
	return fmt.Sprintf("Unknown error #%v [%v]", x.Type, x.Data)
}

func IsInvalidEscapedCharacter(err error) bool {
	x,ok := err.(SplitError)
	if !ok {
		return false
	}
	return x.Type == InvalidEscapedCharacter
}

func IsUnterminatedQuote(err error) bool {
	x,ok := err.(SplitError)
	if !ok {
		return false
	}
	return x.Type == UnterminatedQuote
}
