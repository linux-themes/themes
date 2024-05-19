package main

import (
	"testing"
	"time"
)

func TestA(t *testing.T) {
}

func TestB(t *testing.T) {
	t.Fatal("failed")
}

func TestC(t *testing.T) {
	t.Parallel()
	time.Sleep(time.Second)
	t.Fatal("failed")
}

func TestD(t *testing.T) {
	t.Parallel()
	t.Fatal("failed")
}

func BenchmarkA(b *testing.B) {
}
