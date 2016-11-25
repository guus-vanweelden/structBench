# structBench

## Save Benchmark

```
12:42 ≡ go test -bench .                                                                                                       ᴦ/g/s/g/g/ structBench ≡ MASTER
testing: warning: no tests to run
PASS
BenchmarkSaveStruct_mapStringBool-4      1000000              2082 ns/op             265 B/op          2 allocs/op
BenchmarkSaveStruct_mapBoolString-4      1000000              1098 ns/op             121 B/op          0 allocs/op
BenchmarkSaveStruct_mapStringSlice-4     1000000              1020 ns/op             159 B/op          0 allocs/op
BenchmarkSaveStruct_sliceMapString-4     1000000              1023 ns/op             121 B/op          0 allocs/op
BenchmarkSaveStruct_mapStuct-4           1000000              1056 ns/op             159 B/op          0 allocs/op
BenchmarkLoadStruct_mapStringBool-4     10000000               203 ns/op               0 B/op          0 allocs/op
BenchmarkLoadStruct_mapBoolString-4     10000000               292 ns/op               0 B/op          0 allocs/op
BenchmarkLoadStruct_mapStringSlice-4    10000000               212 ns/op               0 B/op          0 allocs/op
BenchmarkLoadStruct_sliceMapString-4    10000000               243 ns/op               0 B/op          0 allocs/op
BenchmarkLoadStruct_mapStuct-4          10000000               274 ns/op               0 B/op          0 allocs/op
ok      github.com/guus-vanweelden/structBench  128.460s
12:44 ≡      
```
