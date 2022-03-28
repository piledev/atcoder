n, x = map(int,input().split())
ans = -1
alc = 0
for i in range(n):
    v, p = map(int,input().split())
    alc += v*p
    if alc > x*100:
        ans = i+1
        break

print(ans)


