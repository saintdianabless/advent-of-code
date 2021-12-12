using Xunit;
using System;
using MySolutions;

namespace MySolutions.Tests;

public class Day12Test
{
    [Fact]
    public void Part1Test()
    {
        Assert.Equal(10, Day12.Part1("../../../../datas/day12_test1.txt"));
        Assert.Equal(19, Day12.Part1("../../../../datas/day12_test2.txt"));
        Assert.Equal(226, Day12.Part1("../../../../datas/day12_test3.txt"));
    }

    [Fact]
    public void Part2Test()
    {
        Assert.Equal(36, Day12.Part2("../../../../datas/day12_test1.txt"));
    }
}