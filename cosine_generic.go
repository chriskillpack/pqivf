//go:build !cgo

package pqivf

func cosineSimilarityImpl32(a, b []float32) float32 {
	return cosineSimilarityGeneric(a, b)
}

func manhattanDistanceImpl32(a, b []float32) float32 {
	return manhattanDistanceGeneric(a, b)
}
