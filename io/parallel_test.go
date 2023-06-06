package io

import (
	"io"
	"os"
	"strings"
	"testing"
)

func TestParallelWrite1(t *testing.T) {
	src := strings.NewReader("aaa\n")
	out := os.Stdout
	out1 := os.Stdout
	dest := [2]io.Writer{
		out,
		out1,
	}
	err := ParallelWrite(src, dest)
	if err != nil {
		t.Errorf("ParallelWrite err=%v", err)
	}
}

func BenchmarkParallelWrite1(b *testing.B) {
	src := strings.NewReader("aaa\n")
	out := os.Stdout
	out1 := os.Stdout
	dest := [2]io.Writer{
		out,
		out1,
	}
	for n := 0; n < b.N; n++ {
		_ = ParallelWrite(src, dest)
	}
}
