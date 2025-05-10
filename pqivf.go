package pqivf

// CosineSimilarity computes the dot-product between the two equal length
// vectors a & b. The return value is normalized to the range [-1, 1].
// This function will panic if the vectors are unequal length.
func CosineSimilarity(a, b []float32) float32 {
	if len(a) != len(b) {
		panic("Input vectors are unequal lengths")
	}

	if len(a) == 0 {
		return 0
	}

	return cosineSimilarityImpl32(a, b)
}

func ManhattanDistance(a, b []float32) float32 {
	if len(a) != len(b) {
		panic("Input vectors are unequal lengths")
	}

	if len(a) == 0 {
		return 0
	}

	return manhattanDistanceImpl32(a, b)
}
