package service

import (
	"errors"
	"fmt"
)

var ErrInvalidStatType = errors.New("invalid stat type")

type ErrCollectStats struct {
	Err error
}

func (e *ErrCollectStats) Unwrap() error {
	return e.Err
}

func (e *ErrCollectStats) Error() string {
	return fmt.Sprintf("error collecting stats: %v", e.Err)
}
