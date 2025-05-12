#ifdef BUILD_TESTS
    #define DOCTEST_CONFIG_IMPLEMENT_WITH_MAIN
#else
    #define DOCTEST_CONFIG_IMPLEMENT
#endif

#include "doctest.h"
#include "pqivf_neon.h"

TEST_CASE("cosine similarity") {
    SUBCASE("identical vectors") {
        float a[] = {1, 2, 3, 4},
            b[] = {1, 2, 3, 4};
        float res = cosine_similarity_f32_neon(a, b, 4);
        CHECK(res == doctest::Approx(1));
    }

    SUBCASE("orthogonal vectors") {
        float a[] = {1, 0, 0, 0},
            b[] = {0, 1, 0, 0};
        float res = cosine_similarity_f32_neon(a, b, 4);
        CHECK(res == doctest::Approx(0));
    }

    SUBCASE("opposite vectors") {
        float a[] = {1, 2, 3, 4},
            b[] = {-1, -2, -3, -4};
        float res = cosine_similarity_f32_neon(a, b, 4);
        CHECK(res == doctest::Approx(-1));
    }

    SUBCASE("zero vector") {
        float a[] = {1, 2, 3, 4},
            b[] = {0, 0, 0, 0};
        float res = cosine_similarity_f32_neon(a, b, 4);
        CHECK(res == doctest::Approx(0));
    }
}

TEST_CASE("manhattan distance") {
    SUBCASE("") {
        
    }
}