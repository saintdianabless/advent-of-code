#include "utils.h"
#include "solutions/day11.h"

#include <gtest/gtest.h>

TEST(Day11Test, Part1) {
  EXPECT_EQ(D11Part1("../data/day11_test.txt"), 1656l);
}

TEST(Day11Test, Part2) {
  EXPECT_EQ(D11Part2("../data/day11_test.txt"), 195l);
}