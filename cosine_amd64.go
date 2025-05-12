//go:build cgo

package pqivf

// #cgo CFLAGS: -I${SRCDIR}/internal/cfiles/amd64
// #cgo CXXFLAGS:-std=c++11 -I${SRCDIR}/internal/cfiles/amd64
// #include "pqivf_avx.h"
import "C"

import (
	"unsafe"

	_ "github.com/chriskillpack/pqivf/internal/cfiles"

	"golang.org/x/sys/cpu"
)

func cosineSimilarityImpl32(a, b []float32) float32 {
	if !cpu.X86.HasAVX {
		return cosineSimilarityGeneric(a, b)
	}

	return 0.0 // TODO
}

func manhattanDistanceImpl32(a, b []float32) float32 {
	if !cpu.X86.HasAVX {
		return manhattanDistanceGeneric(a, b)
	}

	return float32(C.manhattan_distance_f32_avx(
		(*C.float)(unsafe.Pointer(&a[0])),
		(*C.float)(unsafe.Pointer(&b[0])),
		(C.size_t)(len(a))))
}
