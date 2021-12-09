package com.ybb.aoc.y2021;

import com.ybb.aoc.Utils;

import static org.junit.Assert.assertEquals;
import org.junit.Test;

public class Day09Test {
    @Test
    public void TestPart1() throws Exception {
        var lines = Utils.readResource("y2021/Day09.txt");

        assertEquals(15, Day09.part1(lines));
    }

    @Test
    public void TestPart2() throws Exception {
        assertEquals(1134, Day09.part2(Utils.readResource("y2021/Day09.txt")));
    }
}
