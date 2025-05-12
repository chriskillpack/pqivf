#include "pqivf_neon.h"
#include <math.h>

#ifdef __ARM_NEON
#include <arm_neon.h>
#endif

// TODO - compile error if compiling for a non 64-bit ARM architecture
#ifndef __aarch64__
    #error "Only ARM64 NEON is supported"
#endif

float manhattan_distance_f32_neon(const float* a, const float* b, size_t len) {
    float sum = 0.0f;
    size_t i = 0;

    // Loop while there are at least 4 elements
    for (; i+3 < len ; i+=4) {
        // Load 4 consecutive elements
        float32x4_t va = vld1q_f32(a + i);
        float32x4_t vb = vld1q_f32(b + i);

        // Calculate the absolute difference between the two vectors
        // TODO - could this be replaced by vabdq_f32?
        float32x4_t diff = vsubq_f32(va, vb);
        float32x4_t abs_diff = vabsq_f32(diff);

        sum += vaddvq_f32(abs_diff);
    }

    // It's possible to use NEON intrinsics to handle a 2 element case, the
    // above loop handles 4 element cases. For simplicity do the remaining <=3
    // elements using scalar.
    for (; i<len; i++) {
        sum += fabsf(a[i]-b[i]);
    }

    return sum;
}

float cosine_similarity_f32_neon(const float* a, const float* b, size_t len) {
    size_t i = 0;

    // Create three 4 lane 32-bit float accumulators initialized to 0
    float32x4_t dot_acc = vdupq_n_f32(0);
    float32x4_t a_sqr_acc = vdupq_n_f32(0);
    float32x4_t b_sqr_acc = vdupq_n_f32(0);

    // Loop while there are at least 4 elements
    for (; i+3 < len ; i+=4) {
        // Load 4 consecutive elements
        float32x4_t av = vld1q_f32(a + i);
        float32x4_t bv = vld1q_f32(b + i);

        // Calculate dot-product of vector components a[i]*b[i]
        dot_acc = vmlaq_f32(dot_acc, av, bv);

        // Accumulate length of a and b vectors
        a_sqr_acc = vmlaq_f32(a_sqr_acc, av, av);
        b_sqr_acc = vmlaq_f32(b_sqr_acc, bv, bv);
    }

    // Extract the components of the accumulator
    float dotp = 0.0f, asqrlen = 0.0f, bsqrlen = 0.0f;
    float32_t dot_arr[4],a_sqr_arr[4],b_sqr_arr[4];
    vst1q_f32(dot_arr, dot_acc);
    vst1q_f32(a_sqr_arr, a_sqr_acc);
    vst1q_f32(b_sqr_arr, b_sqr_acc);
    for (int j=0 ; j<4 ; j++) {
        dotp += dot_arr[j];
        asqrlen += a_sqr_arr[j];
        bsqrlen += b_sqr_arr[j];
    }

    // Handle remaining elements
    for (; i<len ; i++) {
        dotp += a[i]*b[i];
        asqrlen += a[i]*a[i];
        bsqrlen += b[i]*b[i];
    }

    float alen = sqrtf(asqrlen), blen = sqrtf(bsqrlen);
	if (alen < 1e-6 || blen < 1e-6) {
		return 0;
	}

    return dotp / (alen * blen);
}