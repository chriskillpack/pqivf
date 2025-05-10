//go:build arm64 && cgo

package pqivf

// #include "pqivf_neon.h"
import "C"
import "unsafe"

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
