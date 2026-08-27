package deferflow

import "errors"

var ErrClose = errors.New("close failed")

type failingCloser struct{}

func (failingCloser) Close() error {
	return ErrClose
}

func IgnoreCloseError() error {
	closer := failingCloser{}
	defer closer.Close()
	return nil
}

func ReturnCloseError() (err error) {
	closer := failingCloser{}
	defer func() {
		if closeErr := closer.Close(); err == nil {
			err = closeErr
		}
	}()
	return nil
}
