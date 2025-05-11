#ifndef PQIVF_NEON_H
#define PQIVF_NEON_H

#include <stddef.h>

#ifdef __cplusplus
extern "C" {
#endif

// Compute the cosine similarity (aka dot product) between vectors a & b
// using NEON instructions. It is the responsbility of the caller to ensure
// that a & b are both equal length.
float cosine_similarity_f32_neon(const float* a, const float* b, size_t len);

// Compute the Manhattan distance (sum of the absolute difference between
// corresponding vector elements) of vectors a & b) using NEON instructions.
// It is the responsbility of the caller to ensure that a & b are both equal
// length.
float manhattan_distance_f32_neon(const float* a, const float* b, size_t len);

#ifdef __cplusplus
}
#endif

#endif