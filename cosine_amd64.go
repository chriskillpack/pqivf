//go:build cgo

package pqivf

import (
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

	return 0.0 // TODO
}
