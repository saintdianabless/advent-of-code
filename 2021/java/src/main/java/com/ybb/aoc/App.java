package com.ybb.aoc;

import com.ybb.aoc.y2021.Day09;

/**
 * Hello world!
 *
 */
public class App {
    public static void main(String[] args) {
        try {
            var lines = Utils.readFile(args[0]);

            System.out.println(Day09.part2(lines));
        } catch (Exception e) {
            System.out.println(e.toString());
            System.exit(1);
        }
    }
}
