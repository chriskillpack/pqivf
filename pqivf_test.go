package pqivf

import (
	"math"
	"math/rand/v2"
	"testing"
)

const benchSize = 4000

type testCase struct {
	name     string
	a        []float32
	b        []float32
	expected float32
}

// isNear returns true if x is within a tolerance of y
func isNear(x, y float32) bool {
	return math.Abs(float64(x)-float64(y)) < float64(1e-4)
}

func TestCosineSimilarity_UnequalLength(t *testing.T) {
	defer func(t *testing.T) {
		if r := recover(); r == nil {
			t.Errorf("Expected panic")
		}
	}(t)

	CosineSimilarity(make([]float32, 5), make([]float32, 6))
}

func TestCosineSimilarity(t *testing.T) {
	tests := []testCase{
		{
			"Identical vectors",
			[]float32{1, 2, 3, 4},
			[]float32{1, 2, 3, 4},
			1,
		},
		{
			name:     "Orthogonal vectors",
			a:        []float32{1, 0, 0, 0},
			b:        []float32{0, 1, 0, 0},
			expected: 0.0,
		},
		{
			name:     "Opposite vectors",
			a:        []float32{1, 2, 3, 4},
			b:        []float32{-1, -2, -3, -4},
			expected: -1.0,
		},
		{
			name:     "Zero vector",
			a:        []float32{0, 0, 0, 0},
			b:        []float32{1, 2, 3, 4},
			expected: 0.0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got, expected := cosineSimilarityGeneric(tc.a, tc.b), tc.expected; !isNear(got, expected) {
				t.Errorf("Expected dot product of %.4f, got %.4f from generic", tc.expected, got)
			}

			if got, expected := CosineSimilarity(tc.a, tc.b), tc.expected; !isNear(got, expected) {
				t.Errorf("Expected dot product of %.4f, got %.4f from optimized", tc.expected, got)
			}
		})
	}
}

func BenchmarkCosineSimilarity(b *testing.B) {
	vecA := make([]float32, benchSize)
	vecB := make([]float32, benchSize)

	for i := range benchSize {
		vecA[i] = rand.Float32() * 100
		vecB[i] = rand.Float32() * 100
	}

	for b.Loop() {
		CosineSimilarity(vecA, vecB)
	}
}

func BenchmarkCosineSimilarityGeneric(b *testing.B) {
	vecA := make([]float32, benchSize)
	vecB := make([]float32, benchSize)

	for i := range benchSize {
		vecA[i] = rand.Float32() * 100
		vecB[i] = rand.Float32() * 100
	}

	for b.Loop() {
		cosineSimilarityGeneric(vecA, vecB)
	}
}

func TestManhattanDistance_UnequalLength(t *testing.T) {
	defer func(t *testing.T) {
		if r := recover(); r == nil {
			t.Errorf("Expected panic")
		}
	}(t)

	ManhattanDistance(make([]float32, 5), make([]float32, 6))
}

func TestManhattanDistance(t *testing.T) {
	tests := []testCase{
		{
			name:     "Single element",
			a:        []float32{1},
			b:        []float32{2},
			expected: 1,
		},
		{
			name:     "Multiple elements",
			a:        []float32{1, 2, 3, 4},
			b:        []float32{5, 6, 7, 8},
			expected: 16,
		},
		{
			name:     "Negative values",
			a:        []float32{-1, -2, -3, -4},
			b:        []float32{1, 2, 3, 4},
			expected: 20,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got, expected := manhattanDistanceGeneric(tc.a, tc.b), tc.expected; !isNear(got, expected) {
				t.Errorf("Expected dot product of %.4f, got %.4f from generic", tc.expected, got)
			}

			if got, expected := ManhattanDistance(tc.a, tc.b), tc.expected; !isNear(got, expected) {
				t.Errorf("Expected dot product of %.4f, got %.4f from optimized", tc.expected, got)
			}
		})
	}
}

func BenchmarkManhattanDistance(b *testing.B) {
	vecA := make([]float32, benchSize)
	vecB := make([]float32, benchSize)

	for i := range benchSize {
		vecA[i] = rand.Float32() * 100
		vecB[i] = rand.Float32() * 100
	}

	for b.Loop() {
		ManhattanDistance(vecA, vecB)
	}
}

func BenchmarkManhattanDistanceGeneric(b *testing.B) {
	vecA := make([]float32, benchSize)
	vecB := make([]float32, benchSize)

	for i := range benchSize {
		vecA[i] = rand.Float32() * 100
		vecB[i] = rand.Float32() * 100
	}

	for b.Loop() {
		manhattanDistanceGeneric(vecA, vecB)
	}
}
