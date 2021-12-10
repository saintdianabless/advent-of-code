lines = []
with open('q.txt') as rawData:
    lines = rawData.readlines()


def part1(lines):
    mapping1 = {')': '(', ']': '[', '}': '{', '>': '<'}
    mapping2 = {')': 3, ']': 57, '}': 1197, '>': 25137}
    st = []
    sum = 0
    for i in range(len(lines)):
        line = lines[i]
        for j in range(len(line)-1):
            c = line[j]
            if c in mapping1.values():
                st.append(c)
            else:
                if len(st) == 0 or st.pop() != mapping1[c]:
                    sum += mapping2[c]
                    break


def part2(lines):
    mapping1 = {')': '(', ']': '[', '}': '{', '>': '<'}
    incomplete = []
    for _, line in enumerate(lines):
        st = []
        for i, c in enumerate(line):
            if len(st) != 0 and c == '\n':
                incomplete.append(st)
                break
            if c in mapping1.values():
                st.append(c)
            else:
                if len(st) == 0 or st.pop() != mapping1[c]:
                    break
    mapping2 = {'(': 1, '[': 2, '{': 3, '<': 4}
    sums = []
    for _, st in enumerate(incomplete):
        sum = 0
        while len(st) != 0:
            sum = sum * 5 + mapping2[st.pop()]
        sums.append(sum)
    sums.sort()
    
    return sums[int(len(sums) / 2)]


print(part2(lines))
