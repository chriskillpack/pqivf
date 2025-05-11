//go:build arm64 && cgo

package pqivf

// #cgo CFLAGS: -I${SRCDIR}/internal/cfiles
// #cgo CXXFLAGS:-std=c++11 -I${SRCDIR}/internal/cfiles
// #include "pqivf_neon.h"
import "C"

import (
	"unsafe"

	_ "github.com/chriskillpack/pqivf/internal/cfiles"
)

func cosineSimilarityImpl32(a, b []float32) float32 {
	return float32(C.cosine_similarity_f32_neon(
		(*C.float)(unsafe.Pointer(&a[0])),
		(*C.float)(unsafe.Pointer(&b[0])),
		(C.size_t)(len(a))))
}

func manhattanDistanceImpl32(a, b []float32) float32 {
	return float32(C.manhattan_distance_f32_neon(
		(*C.float)(unsafe.Pointer(&a[0])),
		(*C.float)(unsafe.Pointer(&b[0])),
		(C.size_t)(len(a))))
}
