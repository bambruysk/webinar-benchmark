package main

import (
	"math/rand"
	"testing"
)

func BenchmarkFindSlice10(b *testing.B) {
	n := 10
	s := storageSlice{
		make([]int, 0, n),
	}

	for i := 0; i < n; i++ {
		s.data = append(s.data, rand.Intn(n/3))
	}

	found := 0
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		found, _ = s.find(3)
	}
	b.StopTimer()
	b.Log(found)
}

func BenchmarkFindSlice100(b *testing.B) {
	n := 100
	s := storageSlice{
		make([]int, 0, n),
	}

	for i := 0; i < n; i++ {
		s.data = append(s.data, rand.Intn(n/10))
	}

	found := 0
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		found, _ = s.find(3)
	}
	b.StopTimer()
	b.Log(found)
}

func BenchmarkFindSlice1000(b *testing.B) {
	n := 1000
	s := storageSlice{
		make([]int, 0, n),
	}

	for i := 0; i < n; i++ {
		s.data = append(s.data, rand.Intn(n/10))
	}

	found := 0
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		found, _ = s.find(3)
	}
	b.StopTimer()
	b.Log(found)
}

func BenchmarkFindMap10(b *testing.B) {
	s := storageMap{
		make(map[int]struct{}, 10),
	}

	for i := 0; i < 10; i++ {
		s.data[i] = struct{}{}
	}

	found := 0
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		found, _ = s.find(3)
	}
	b.StopTimer()
	b.Log(found)
}

func BenchmarkFindMap100(b *testing.B) {
	s := storageMap{
		make(map[int]struct{}, 100),
	}

	for i := 0; i < 100; i++ {
		s.data[i] = struct{}{}
	}

	found := 0
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		found, _ = s.find(3)
	}
	b.StopTimer()
	b.Log(found)
}

func BenchmarkFindMap1000(b *testing.B) {
	s := storageMap{
		make(map[int]struct{}, 1000),
	}

	for i := 0; i < 1000; i++ {
		s.data[i] = struct{}{}
	}

	found := 0
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		found, _ = s.find(3)
	}
	b.StopTimer()
	b.Log(found)
}

func BenchmarkFindMap10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := storageMap{
			make(map[int]struct{}, 10000),
		}

		for i := 0; i < 10000; i++ {
			s.data[i] = struct{}{}
		}

		found := 0

		found, _ = s.find(3)
		b.StopTimer()
		b.Log(found)
	}
}
