from typing import Counter

def part1(file):
    patterns = {}
    with open('day14.txt', 'r') as f:
        polymer = f.readline().strip()
        print(polymer)
        f.readline()

        for line in f:
            line = line.strip()
            (p, c) = line.split(' -> ')
            patterns[p] = c

        for i in range(10):
            inserted = ""
            for i in range(0, len(polymer) - 1):
                inserted += polymer[i]
                sub = polymer[i:i+2]
                if sub in patterns.keys():
                    inserted += patterns[sub]
            inserted += polymer[-1]
            polymer = inserted
        chars = set(polymer)
        counts = {}
        for c in chars:
            counts[c] = polymer.count(c)
        maximum = max(counts.values())
        minimum = min(counts.values())
        print(maximum - minimum)


def solve(s, patterns):
    c1 = Counter()
    newinserted = Counter(s)
    for i in range(0, len(s) - 1):
        c1[s[i:i+2]] += 1
    
    for i in range(40):
        # print(c1)
        c2 = Counter()
        for c, cnt in c1.most_common():
            if c in patterns.keys():
                newinserted[patterns[c]] += 1
                c2[c[0] + patterns[c]] += cnt
                c2[patterns[c] + c[1]] += cnt
            else:
                c2[c] += c1[c]
        c1 = c2

    c0 = Counter()
    for a, b in c1.most_common():
        c0[a[0]] += b
    c0[s[-1]] += 1

    print(max(c0.values()) - min(c0.values()))
    
def part2(file):
    patterns = {}
    polymer = ""
    with open(file, 'r') as f:
        polymer = f.readline().strip()
        f.readline()
        for line in f:
            line = line.strip()
            (p, c) = line.split(' -> ')
            patterns[p] = c

    solve(polymer, patterns)

part2('day14.txt')
