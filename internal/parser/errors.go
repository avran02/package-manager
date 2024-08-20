package parser

import (
	"errors"
)

var (
	ErrUnknownAction   = errors.New("unknown action")
	ErrUnknownAFileExt = errors.New("unknown file type")
)
