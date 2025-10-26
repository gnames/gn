package gn

import (
	"errors"
)

func PrintErrorMessage(err error) {
	var target Error
	printable := errors.As(err, target)
	if printable {
		target.print()
	}
}

type ErrorCode int

type Error struct {
	Code ErrorCode
	Err  error
	Msg  string
	Vars []any
}

func (e Error) print() {
	print(errorMsgType, e.Msg, e.Vars)
}
