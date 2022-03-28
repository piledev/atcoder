N,K=map(int,input().split())
t = []
for i in range(N):
    t.append(list(map(int,input().split())))
ans = 0    
import itertools
for route in itertools.permutations(range(1,N)):
    nowtime = 0
    nowtime += t[0][route[0]]
    nowcity = route[0]
    for i in range(1,N-1):
        tocity = route[i]
        nowtime += t[nowcity][tocity]
        nowcity = tocity
    nowtime += t[nowcity][0]
    if nowtime == K:
        ans += 1
print(ans)
    
