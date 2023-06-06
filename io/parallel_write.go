package io

import (
	"io"
)

func ParallelWrite(src io.Reader, dest [2]io.Writer) (err error) {
	errChan := make(chan error, 1)

	pr, pw := io.Pipe()
	wr := io.TeeReader(src, pw)
	go func() {
		var _err error
		defer func() {
			pr.CloseWithError(_err)
			errChan <- _err
		}()
		_, _err = io.Copy(dest[0], pr)
	}()
	defer func() {
		pw.Close()
		_err := <-errChan
		_ = _err
	}()

	_, err = io.Copy(dest[1], wr)
	return err
}
