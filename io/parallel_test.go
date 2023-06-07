package io

import (
	"io"
	"os"
	"strings"
	"testing"
)

func TestParallelMultiWrite(t *testing.T) {
	writers := []io.Writer{
		os.Stdout, os.Stdout, os.Stdout,
	}
	if err := ParallelMultiWrite(strings.NewReader("aaa\n"), writers); err != nil {
		t.Errorf("ParallelMultiWrite err=%v", err)
	}
}

func TestParallelWrite1(t *testing.T) {
	src := strings.NewReader("aaa\n")
	dest := [2]io.Writer{
		os.Stdout,
		os.Stdout,
	}
	err := ParallelWrite(src, dest)
	if err != nil {
		t.Errorf("ParallelWrite err=%v", err)
	}
}

func BenchmarkParallelWrite1(b *testing.B) {
	src := strings.NewReader("aaa\n")
	dest := [2]io.Writer{
		io.Discard,
		io.Discard,
	}
	for n := 0; n < b.N; n++ {
		_ = ParallelWrite(src, dest)
	}
}

func BenchmarkParallelMultiWrite(b *testing.B) {
	src := strings.NewReader("aaa\n")
	dests := []io.Writer{
		io.Discard, io.Discard,
	}
	for n := 0; n < b.N; n++ {
		_ = ParallelMultiWrite(src, dests)
	}
}
