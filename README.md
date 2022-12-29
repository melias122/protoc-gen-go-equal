# protoc-gen-go-equal

protoc-gen-go-equal is a protobuf plugin that generates `Equal` methods

### Installation
`go get github.com/melias122/protoc-gen-go-equal@latest`

### Benchmark 
`proto.Equal` vs generated `Equal`
```
name                                  old time/op    new time/op    delta
EqualWithSmallEmpty-16                   885ns ± 1%       3ns ±31%   -99.65%  (p=0.016 n=4+5)
EqualWithIdenticalPtrEmpty-16           7.96ns ±37%    2.51ns ±33%   -68.44%  (p=0.008 n=5+5)
EqualWithLargeEmpty-16                  8.10µs ±26%    0.48µs ±37%   -94.11%  (p=0.008 n=5+5)
EqualWithDeeplyNestedEqual-16            184µs ±23%       7µs ± 2%   -95.93%  (p=0.016 n=5+4)
EqualWithDeeplyNestedDifferent-16       58.8µs ± 4%     1.2µs ± 2%   -97.98%  (p=0.016 n=5+4)
EqualWithDeeplyNestedIdenticalPtr-16    5.07ns ± 5%    2.27ns ± 1%   -55.12%  (p=0.016 n=5+4)

name                                  old alloc/op   new alloc/op   delta
EqualWithSmallEmpty-16                   88.0B ± 0%      0.0B       -100.00%  (p=0.008 n=5+5)
EqualWithIdenticalPtrEmpty-16            0.00B          0.00B           ~     (all equal)
EqualWithLargeEmpty-16                   88.0B ± 0%      0.0B       -100.00%  (p=0.008 n=5+5)
EqualWithDeeplyNestedEqual-16           3.52kB ± 0%    0.00kB       -100.00%  (p=0.008 n=5+5)
EqualWithDeeplyNestedDifferent-16       2.58kB ± 0%    0.00kB       -100.00%  (p=0.008 n=5+5)
EqualWithDeeplyNestedIdenticalPtr-16     0.00B          0.00B           ~     (all equal)

name                                  old allocs/op  new allocs/op  delta
EqualWithSmallEmpty-16                    5.00 ± 0%      0.00       -100.00%  (p=0.008 n=5+5)
EqualWithIdenticalPtrEmpty-16             0.00           0.00           ~     (all equal)
EqualWithLargeEmpty-16                    5.00 ± 0%      0.00       -100.00%  (p=0.008 n=5+5)
EqualWithDeeplyNestedEqual-16              200 ± 0%         0       -100.00%  (p=0.008 n=5+5)
EqualWithDeeplyNestedDifferent-16          122 ± 0%         0       -100.00%  (p=0.008 n=5+5)
EqualWithDeeplyNestedIdenticalPtr-16      0.00           0.00           ~     (all equal)
```

### Limitations
Use only when speed and efficiency is a concern, otherwise use `proto.Equal`.
In some cases as e.g. some known types `Equal` will fallback to `proto.Equal`.
Unknown fields and extensions are not supported.
