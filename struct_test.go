package main

import "testing"

func benchmarkInsertStruct(s saverLoader, b *testing.B) {
	d := generateItems(b.N)
	b.ResetTimer()
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		s.Save(d[n])
	}
}

func benchmarkLoadStruct(s saverLoader, b *testing.B) {
	d := generateItems(b.N)
	k := make([]key, len(d))
	for n := 0; n < b.N; n++ {
		if n%2 == 0 {
			s.Save(d[n])
		}
		k[n] = d[n].Key()
	}
	b.ResetTimer()
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		s.Load(k[n])
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

func BenchmarkLoadStruct_mapStringBool(b *testing.B) {
	m := make(mapStringBool)
	benchmarkLoadStruct(m, b)
}

func BenchmarkLoadStruct_mapBoolString(b *testing.B) {
	m := make(mapBoolString)
	benchmarkLoadStruct(m, b)
}

func BenchmarkLoadStruct_mapStringSlice(b *testing.B) {
	m := make(mapStringSlice)
	benchmarkLoadStruct(m, b)
}

func BenchmarkLoadStruct_sliceMapString(b *testing.B) {
	m := sliceMapString{}
	m[0] = make(map[string]*item)
	m[1] = make(map[string]*item)
	benchmarkLoadStruct(m, b)
}

func BenchmarkLoadStruct_mapStuct(b *testing.B) {
	m := make(mapStruct)
	benchmarkLoadStruct(m, b)
}
