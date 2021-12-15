import collections
import functools
import math

class impossible(Exception):
    def __init__(self, *args: object) -> None:
        super().__init__(*args)

def H(p, r, c):
    x = c - p[1]
    y = r - p[0]
    return math.sqrt(x * x + y * y)

def get_neighbors(n, grids, r, c):
    ret = []
    if n[0]+1 != r:
        b = (n[0] + 1, n[1])
        ret.append((b, grids[b[0]][b[1]]))
    if n[0]-1 >= 0:
        b = (n[0] - 1, n[1])
        ret.append((b, grids[b[0]][b[1]]))
    if n[1]+1 != c:
        b = (n[0], n[1] + 1)
        ret.append((b, grids[b[0]][b[1]]))
    if n[1]-1 >= 0:
        b = (n[0], n[1] - 1)
        ret.append((b, grids[b[0]][b[1]]))

    return ret

def a_star(grids):
    open_set = set()
    open_set.add((0, 0))
    closed_set = set()

    parents = {}
    parents[(0, 0)] = (0, 0)

    G = {}
    G[(0, 0)] = 0

    r = len(grids)
    c = len(grids[0])

    end = (r-1, c-1)

    while len(open_set) > 0:
        N = None

        for v in open_set:
            if N == None or G[v] + H(v, r, c) < G[N] + H(v, r, c):
                N = v

        if N == end or get_neighbors(N, grids, r, c) == []:
            pass
        else:
            for n, weight in get_neighbors(N, grids, r, c):
                if n not in open_set and n not in closed_set:
                    open_set.add(n)
                    parents[n] = N
                    G[n] = G[N] + weight
                else:
                    if G[n] > G[N] + weight:
                        G[n] = G[N] + weight
                        parents[n] = N

                        if n in closed_set:
                            closed_set.remove(n)
                            open_set.add(n)
        if N == None:
            raise impossible
        if N == end:
            path = [end]
            node = end
            while node != (0, 0):
                path.append(parents[node])
                node = parents[node]
            return path
        
        open_set.remove(N)
        closed_set.add(N)


def part1(path):
    grids = []
    with open(path, 'r') as f:
        for line in f.readlines():
            l = []
            for c in range(len(line)-1):
                l.append(ord(line[c]) - ord('0'))
            grids.append(l)
    path = a_star(grids=grids)
    print(path)
    for i in range(len(grids)):
        for j in range(len(grids[0])):
            if (i, j) in path:
                print('.', end='')
            else:
                print(grids[i][j], end='')
        print()
    cost = 0
    for n in path:
        cost += grids[n[0]][n[1]]
    cost -= grids[0][0]
    print(cost)

def part2(path, N):
    grids = []
    with open(path, 'r') as f:
        for line in f.readlines():
            l = []
            for c in range(len(line)-1):
                l.append(ord(line[c]) - ord('0'))
            size = len(l)
            for i in range(1, N):
                for j in range(size):
                    e = l[(i-1)*size+j] + 1
                    if e == 10:
                        l.append(1)
                    else:
                        l.append(e)
            grids.append(l)
        size = len(grids)
        for i in range(1, N):
            for j in range(size):
                last = grids[(i-1)*size + j]
                me = []
                for e in last:
                    if e + 1 == 10:
                        me.append(1)
                    else:
                        me.append(e+1)
                grids.append(me)
    path = a_star(grids=grids)
    print(path)
    for i in range(len(grids)):
        for j in range(len(grids[0])):
            if (i, j) in path:
                print('.', end='')
            else:
                print(grids[i][j], end='')
        print()
    cost = 0
    for n in path:
        cost += grids[n[0]][n[1]]
    cost -= grids[0][0]
    print(cost)

part1('./d15.txt')
part2('./d15.txt', 5)
