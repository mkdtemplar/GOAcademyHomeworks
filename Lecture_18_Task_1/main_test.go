package main

import (
	"testing"
)

func Benchmark100GoPrimesAndSleepWith0MsSleep(b *testing.B) {
	for i := 0; i < b.N; i++ {
		goPrimesAndSleep(100, 0)
	}
}

func Benchmark100PrimesWith5MSSleep(b *testing.B) {
	for i := 0; i < b.N; i++ {
		goPrimesAndSleep(100, 5)
	}
}
