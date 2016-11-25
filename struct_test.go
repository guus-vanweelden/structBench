package main

import "testing"

func benchmarkInsertStruct(s saverLoader, b *testing.B) {
	d := generateItems(b.N)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		s.Save(d[n])
	}
}

func BenchmarkSaveStruct_mapStringBool(b *testing.B) {
	m := make(mapStringBool)
	benchmarkInsertStruct(m, b)
}

func BenchmarkSaveStruct_mapBoolString(b *testing.B) {
	m := make(mapBoolString)
	benchmarkInsertStruct(m, b)
}

func BenchmarkSaveStruct_mapStringSlice(b *testing.B) {
	m := make(mapStringSlice)
	benchmarkInsertStruct(m, b)
}

func BenchmarkSaveStruct_sliceMapString(b *testing.B) {
	m := sliceMapString{}
	m[0] = make(map[string]*item)
	m[1] = make(map[string]*item)
	benchmarkInsertStruct(m, b)
}

func BenchmarkSaveStruct_mapStuct(b *testing.B) {
	m := make(mapStruct)
	benchmarkInsertStruct(m, b)
}
