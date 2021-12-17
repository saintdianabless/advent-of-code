import re
import time
from numba import jit

p = ''
with open('17.txt', 'r') as f:
    p = f.readline()
    
r = re.compile(r'^.*=([-\d]+)..([-\d]+).*=([-\d]+)..([-\d]+)$')
m = r.match(p)
pos = list(m.group(1, 2, 3, 4))
for i in range(4):
    pos[i] = int(pos[i])
   
@jit(nopython=True) 
def throw_it(a, b, c, d):
    cnt = 0
    max_y = 0
    for v_x in range(b + 1):
        for v_y in range(c, 1000):
            x = 0
            y = 0
            vx = v_x
            vy = v_y
            highest = -1
            ok = False
            for _ in range(500):
                x += vx
                y += vy
                if x > b or y < c or (vx == 0 and y < c):
                    break
                highest = max(highest, y)
                if a <= x <= b and c <= y <= d:
                    ok = True
                if vx != 0:
                    vx -= 1
                vy -= 1
            if ok:
                cnt += 1
                if max_y < highest:
                    max_y = highest
    return max_y, cnt

s = time.time()
print(throw_it(*pos))
print('{:.4}s'.format(time.time() - s))
