package main

import (
	"fmt"
	"testing"
	"time"
)

func Benchmark100PrimesWith0MSSleep(b *testing.B) {
	for i := 0; i < b.N; i++ {
		primesAndSleep(100, 0)
	}
}

func BenchmarkBenchmark100PrimesWith5MSSleep(b *testing.B) {
	for i := 0; i < b.N; i++ {
		primesAndSleep(100, 5)
	}
}

func BenchmarkBenchmark100PrimesWith10MSSleep(b *testing.B) {
	for i := 0; i < b.N; i++ {
		primesAndSleep(100, 10)
	}
}

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

func Benchmark100GoPrimesWith10MSSleep(b *testing.B) {
	for i := 0; i < b.N; i++ {
		goPrimesAndSleep(100, 10)
	}
}

func BenchmarkGoPrimesWithSleep(b *testing.B) {
	sleepTimes := []time.Duration{time.Millisecond * 0, time.Millisecond * 5, time.Millisecond * 10}
	for _, st := range sleepTimes {
		b.Run(fmt.Sprintf("Benchmark-sleep-time-%d", st), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				goPrimesAndSleep(100, st)
			}
		})
	}

}

func Benchmark100PrimesSleep(b *testing.B) {
	sleepTimes := []time.Duration{time.Millisecond * 0, time.Millisecond * 5, time.Millisecond * 10}
	for _, st := range sleepTimes {
		b.Run(fmt.Sprintf("Benchmark-sleep-time-%d", st), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				primesAndSleep(100, st)
			}
		})
	}
}