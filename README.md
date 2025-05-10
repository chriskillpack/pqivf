# PQIVF library for Go

This Go package implements Manhattan Distance and Cosine Similarity for 32-bit floats. Under the hood it will use SIMD instructions to accelerate the computations when running on ARM64. Alternatively if you build for a non-ARM64 platform or with CGO disabled it will use pure Go implementations of the functions. I hope to add support for AVX (AMD64 SIMD instructions) at a later date.

The NEON intrinsic versions are significantly faster:
```
$ go test -bench=.
...
cpu: Apple M1
BenchmarkCosineSimilarity-8               895324              1330 ns/op
BenchmarkCosineSimilarityGeneric-8        241170              4969 ns/op
BenchmarkManhattanDistance-8             1223034               980.6 ns/op
BenchmarkManhattanDistanceGeneric-8       318074              3742 ns/op
```

`BenchmarkCosineSimilarityGeneric` and `BenchmarkManhattanDistanceGeneric` benchmark the pure Go implementations.

## How to use

Use `CosineSimilarity` to compute the cosine similarity between two vectors, and `ManhattanDistance` to compute the '[Manhattan distance](https://simple.wikipedia.org/wiki/Manhattan_distance)' between them. The vectors are represented as `[]float32`, and the routines will panic if the input slices are different lengths.

None of the functions handle NaN's, so garbage in will give you garbage output.

## Testing the C code

The optimized versions that use SIMD are implemented in C files via intrinsics. They come with their own unit tests which can run via the Makefile. The unit tests use [doctest](https://github.com/doctest/doctest). Currently the Makefile assumes you have a recent version of `clang++` with support for NEON intrinsics.

## TODO

* amd64 AVX intrinsics acceleration of Manhattan Distance and Cosine Similarity.
* float64 support
