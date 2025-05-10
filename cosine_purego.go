package pqivf

import (
	"math"
)

func cosineSimilarityGeneric(a, b []float32) float32 {
	var sum float32

	var magA, magB float32
	for i := range len(a) {
		sum += a[i] * b[i]
		magA += a[i] * a[i]
		magB += b[i] * b[i]
	}

	if float64(magA) < 1e-6 || float64(magB) < 1e-6 {
		return 0
	}

	return sum / float32(math.Sqrt(float64(magA))*math.Sqrt(float64(magB)))
}

func manhattanDistanceGeneric(a, b []float32) float32 {
	var sum float32

	for i := range len(a) {
		sum += float32(math.Abs(float64(a[i] - b[i])))
	}

	return sum
}
