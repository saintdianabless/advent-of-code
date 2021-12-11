cmake_minimum_required(VERSION 3.19)

project(aoc2021)

set(CMAKE_CXX_STANDARD 14)

enable_testing()

find_package(folly CONFIG REQUIRED)
find_package(GTest CONFIG REQUIRED)

add_executable(
  howto
  howto.cpp
)

target_link_libraries(
  howto
  Folly::folly
)

file(GLOB solutions_SRC solutions/*.h)
message(${solutions_SRC})

add_executable(
  runner
  runner.cpp
  utils.h
  ${solutions_SRC} 
)

target_link_libraries(
  runner
  Folly::folly
)

add_executable(
  all_tests
  all_tests.cpp
  utils.h
  ${solutions_SRC}
)
target_link_libraries(
  all_tests
  Folly::folly
  GTest::gtest_main
)

include(GoogleTest)
gtest_discover_tests(all_tests)