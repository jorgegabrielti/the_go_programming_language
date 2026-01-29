package main

import (
	"strings"
	"testing"
)

// Simula os argumentos de linha de comando
var testArgs = []string{
	"programa",
	"Hello",
	"World",
	"Go",
	"is",
	"awesome",
	"and",
	"very",
	"fast",
	"language",
}

// Benchmark do método echo1 (loop tradicional com índice)
func BenchmarkEcho1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var s, sep string
		for j := 1; j < len(testArgs); j++ {
			s += sep + testArgs[j]
			sep = " "
		}
		_ = s // Evita otimização do compilador
	}
}

// Benchmark do método echo2 (range)
func BenchmarkEcho2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s, sep := "", ""
		for _, arg := range testArgs[1:] {
			s += sep + arg
			sep = " "
		}
		_ = s // Evita otimização do compilador
	}
}

// Benchmark do método echo3 (strings.Join)
func BenchmarkEcho3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := strings.Join(testArgs[1:], " ")
		_ = s // Evita otimização do compilador
	}
}

// Teste com MUITOS argumentos (100 argumentos)
var manyArgs = make([]string, 101)

func init() {
	manyArgs[0] = "programa"
	for i := 1; i <= 100; i++ {
		manyArgs[i] = "arg"
	}
}

func BenchmarkEcho1_ManyArgs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var s, sep string
		for j := 1; j < len(manyArgs); j++ {
			s += sep + manyArgs[j]
			sep = " "
		}
		_ = s
	}
}

func BenchmarkEcho2_ManyArgs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s, sep := "", ""
		for _, arg := range manyArgs[1:] {
			s += sep + arg
			sep = " "
		}
		_ = s
	}
}

func BenchmarkEcho3_ManyArgs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := strings.Join(manyArgs[1:], " ")
		_ = s
	}
}
