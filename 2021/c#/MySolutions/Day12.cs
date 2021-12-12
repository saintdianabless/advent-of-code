namespace MySolutions;

class Cave
{
    public String _tag { get; set; }
    public bool _small { get; set; }
    public int _visited { get; set; }
    public List<Cave> _adjs { get; set; }

    public Cave(bool small, String tag)
    {
        _tag = tag;
        _small = small;
        _adjs = new List<Cave>();
    }
}

public class Day12
{

    private static Dictionary<String, Cave> initGraph(String filename)
    {
        var graph = new Dictionary<String, Cave>();
        var start = new Cave(true, "start") { _visited = 100 };
        var end = new Cave(true, "end");
        graph["start"] = start;
        graph["end"] = end;

        using (var stream = File.OpenText(filename))
        {
            while (stream.ReadLine() is String s)
            {
                var splited = s.Split('-');
                if (!graph.ContainsKey(splited[0]))
                {
                    graph[splited[0]] = new Cave(Char.IsLower(splited[0][0]), splited[0]);
                }
                if (!graph.ContainsKey(splited[1]))
                {
                    graph[splited[1]] = new Cave(Char.IsLower(splited[1][0]), splited[1]);
                }
                graph[splited[0]]._adjs.Add(graph[splited[1]]);
                graph[splited[1]]._adjs.Add(graph[splited[0]]);
            }
        }

        return graph;
    }

    private static void dfs(Cave from, Dictionary<String, Cave> graph, ref int count)
    {
        foreach (var adj in from._adjs)
        {
            if (adj == graph["end"])
            {
                count += 1;
                continue;
            }
            if (adj._small && adj._visited != 0)
            {
                continue;
            }
            adj._visited += 1;
            dfs(adj, graph, ref count);
            adj._visited -= 1;
        }
    }
    public static int Part1(String filename)
    {
        var graph = initGraph(filename);

        int count = 0;
        dfs(graph["start"], graph, ref count);
        return count;
    }

    private static void dfs(Cave from, Dictionary<String, Cave> graph, ref int count, bool alreadyTwice)
    {
        foreach (var adj in from._adjs)
        {
            if (adj == graph["end"])
            {
                count += 1;
                continue;
            }
            if (alreadyTwice)
            {
                if (adj._small && adj._visited != 0)
                {
                    continue;
                }
                adj._visited += 1;
                dfs(adj, graph, ref count, true);
                adj._visited -= 1;
            } else {
                if(adj._small && adj._visited == 1) {
                    adj._visited += 1;
                    dfs(adj, graph, ref count, true);
                    adj._visited -= 1;
                } else if(!adj._small || adj._visited == 0) {
                    adj._visited += 1;
                    dfs(adj, graph, ref count, false);
                    adj._visited -= 1;
                }
            }
        }
    }

    public static int Part2(String filename)
    {
        var graph = initGraph(filename);

        int count = 0;
        dfs(graph["start"], graph, ref count, false);
        return count;
    }
}
