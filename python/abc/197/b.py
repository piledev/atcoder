h, w, x, y = map(int, input().split())
grid = [str(input()) for _ in range(h)] 
x -= 1
y -= 1
ans = 1  

for i in range(1,h):
    if 0 <= x-i < h :
        if grid[x-i][y] == "#":
            break
        ans += 1

for i in range(1,h):
    if 0 <= x+i < h :
        if grid[x+i][y] == "#":
            break
        ans += 1

for j in range(1,w):
    if 0 <= y-j < w:
        if grid[x][y-j] == "#":
            break
        ans += 1

for j in range(1,w):
    if 0 <= y+j < w:
        if grid[x][y+j] == "#":
            break
        ans += 1

print(ans)
