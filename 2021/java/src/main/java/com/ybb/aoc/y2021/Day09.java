package com.ybb.aoc.y2021;

import java.util.ArrayList;
import java.util.Collections;
import java.util.List;

public final class Day09 {
    private Day09() {
    }

    private static boolean isLowest(List<String> mat, int r, int c) {
        var me = mat.get(r).charAt(c);
        if ((r - 1 < 0 || mat.get(r - 1).charAt(c) > me) &&
            (c - 1 < 0 || mat.get(r).charAt(c - 1) > me) &&
            (r + 1 >= mat.size() || mat.get(r + 1).charAt(c) > me) &&
            (c + 1 >= mat.get(0).length() || mat.get(r).charAt(c + 1) > me)) {
            return true;
        }

        return false;
    }

    public static int part1(List<String> lines) {
        var rows = lines.size();
        var columns = lines.get(0).length();

        int sum = 0;
        for (int i = 0; i < rows; i++) {
            for (int j = 0; j < columns; j++) {
                if(isLowest(lines, i, j)) {
                    sum += lines.get(i).charAt(j) - '0' + 1;
                }
            }
        }

        return sum;
    }

    public static int part2(List<String> lines) {
        var rows = lines.size();
        var columns = lines.get(0).length();
        boolean[][] grids = new boolean[rows][columns];
        for(var i = 0; i < rows; i++) {
            for(var j = 0; j < columns; j++) {
                if(lines.get(i).charAt(j) == '9') {
                    grids[i][j] = true;
                }
            }
        }

        var areas = new ArrayList<Integer>();
        for(var i = 0; i < rows; i++) {
            for(var j = 0; j < columns; j++) {
                if(!grids[i][j]) {
                    int[] area = {0};
                    dfs(area, i, j, grids);
                    areas.add(area[0]);
                }
            }
        }
        areas.sort(Collections.reverseOrder());
        var res = 1;
        for (var i = 0; i < 3; i++) {
            res *= areas.get(i);
        }
        return res;
    }

    public static void dfs(int[] area, int i, int j, boolean[][] grids){
        if(i < 0 || i >= grids.length || j < 0 || j >= grids[0].length || grids[i][j]) return;
        area[0] += 1;
        grids[i][j] = true;

        dfs(area, i-1, j, grids);
        dfs(area, i+1, j, grids);
        dfs(area, i, j-1, grids);
        dfs(area, i, j+1, grids);
    }
}
