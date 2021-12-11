#pragma once
#include "../utils.h"

#include <stdio.h>
#include <vector>
#include <functional>

using folly::fbvector;

void step1(fbvector<fbvector<int>> &grids)
{
    for (auto &r : grids)
    {
        for (auto &e : r)
        {
            ++e;
        }
    }
}

void infect(fbvector<fbvector<int>> &grids, int r, int c)
{
    int rows = grids.size();
    int columns = grids[0].size();
    int end = columns * rows - 1;
    int pos = columns * r + c;

    for (int i = -1; i <= 1; ++i)
    {
        for (int j = -1; j <= 1; ++j)
        {
            if (i == 0 && j == 0)
            {
                continue;
            }
            int pos_r = r + i;
            int pos_c = c + j;
            if (pos_r >= 0 && pos_r < rows && pos_c >= 0 && pos_c < columns)
            {
                int &element = grids[pos_r][pos_c];
                if (element == 11)
                {
                    continue;
                }
                ++element;
                if (element >= 10)
                {
                    element = 11;
                    infect(grids, pos_r, pos_c);
                }
            }
        }
    }
}

void step2(fbvector<fbvector<int>> &grids)
{
    int rows = grids.size();
    int columns = grids[0].size();
    for (int i = 0; i < rows; ++i)
    {
        for (int j = 0; j < columns; ++j)
        {
            if (grids[i][j] == 11)
            {
                continue;
            }
            if (grids[i][j] == 10)
            {
                grids[i][j] = 11;
                infect(grids, i, j);
            }
        }
    }
}

long step3(fbvector<fbvector<int>> &grids)
{
    long cnt = 0;
    for (auto &r : grids)
    {
        for (auto &c : r)
        {
            if (c == 11 || c == 10)
            {
                ++cnt;
                c = 0;
            }
        }
    }
    return cnt;
}

void printGrids(const fbvector<fbvector<int>> &grids)
{
    for (const auto &r : grids)
    {
        for (auto c : r)
        {
            printf("%2d ", c);
        }
        printf("\n");
    }
    printf("-------------------------------------\n");
}

fbvector<fbvector<int>> getGrids(folly::StringPiece filename)
{
    auto lines = readLines(filename);

    int rows = lines.size();
    int columns = lines[0].size();
    auto grids = fbvector<fbvector<int>>(rows, fbvector<int>(columns, 0));
    for (int i = 0; i < rows; ++i)
    {
        for (int j = 0; j < columns; ++j)
        {
            grids[i][j] = lines[i][j] - '0';
        }
    }
    return grids;
}

long D11Part1(folly::StringPiece filename)
{
    auto grids = getGrids(filename);

    long flashed = 0;
    for (int i = 0; i < 100; ++i)
    {
        step1(grids);
        step2(grids);
        flashed += step3(grids);
    }

    return flashed;
}

void step(fbvector<fbvector<int>> &grids, bool &flag)
{
    auto sum = grids.size() * grids[0].size();
    long cnt = 0;
    for (auto &r : grids)
    {
        for (auto &c : r)
        {
            if (c == 11 || c == 10)
            {
                ++cnt;
                c = 0;
            }
        }
    }
    flag = (cnt == sum);
}

long D11Part2(folly::StringPiece filename)
{
    auto grids = getGrids(filename);

    bool flag = false;
    int steps = 0;
    while (!flag)
    {
        ++steps;
        step1(grids);
        step2(grids);
        step(grids, flag);
    }

    return steps;
}