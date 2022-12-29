# protoc-gen-go-equal

protoc-gen-go-equal is a protobuf plugin that generates `Equal` methods

### Installation
`go get github.com/melias122/protoc-gen-go-equal@latest`

### Benchmark 
`proto.Equal`
```
goos: linux
goarch: amd64
pkg: github.com/melias122/protoc-gen-go-equal
cpu: AMD Ryzen 7 PRO 4750U with Radeon Graphics
BenchmarkProtoEqualWithSmallEmpty-16                      1330669                876.6 ns/op	      88 B/op	       5 allocs/op
BenchmarkProtoEqualWithIdenticalPtrEmpty-16             140363698	         8.319 ns/op	       0 B/op	       0 allocs/op
BenchmarkProtoEqualWithLargeEmpty-16                  	   167653	          7033 ns/op	      88 B/op	       5 allocs/op
BenchmarkProtoEqualWithDeeplyNestedEqual-16           	     6852	        175633 ns/op	    3520 B/op	     200 allocs/op
BenchmarkProtoEqualWithDeeplyNestedDifferent-16       	    18258	         60185 ns/op	    2584 B/op	     122 allocs/op
BenchmarkProtoEqualWithDeeplyNestedIdenticalPtr-16    	206724327	         5.299 ns/op	       0 B/op	       0 allocs/op
```

generated `Equal`
```
goos: linux
goarch: amd64
pkg: github.com/melias122/protoc-gen-go-equal
cpu: AMD Ryzen 7 PRO 4750U with Radeon Graphics
BenchmarkEqualWithSmallEmpty-16                         301572816	         3.737 ns/op	       0 B/op	       0 allocs/op
BenchmarkEqualWithIdenticalPtrEmpty-16                  334314547	         3.041 ns/op	       0 B/op	       0 allocs/op
BenchmarkEqualWithLargeEmpty-16                       	  2811765	         383.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkEqualWithDeeplyNestedEqual-16                	   169820	          7399 ns/op	       0 B/op	       0 allocs/op
BenchmarkEqualWithDeeplyNestedDifferent-16            	   957091	          1166 ns/op	       0 B/op	       0 allocs/op
BenchmarkEqualWithDeeplyNestedIdenticalPtr-16         	595409228	         2.009 ns/op	       0 B/op	       0 allocs/op
```

### Limitations
Use only when speed and efficiency is a concern, otherwise use `proto.Equal`.
In some cases as e.g. some known types `Equal` will fallback to `proto.Equal`.
Unknown fields and extensions are not supported.
