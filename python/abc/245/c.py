n,k=map(int,input().split())
a=list(map(int,input().split()))
b=list(map(int,input().split()))

visited_a=[False]*n
visited_b=[False]*n

import sys
sys.setrecursionlimit(250000)

def dfs(dep,val):
    if dep==n-1:
        print('Yes')
        exit()

    if not visited_a[dep+1] and abs(a[dep+1]-val)<=k:
        visited_a[dep+1]=True
        dfs(dep+1,a[dep+1])
    if not visited_b[dep+1] and abs(b[dep+1]-val)<=k:
        visited_b[dep+1]=True
        dfs(dep+1,b[dep+1])
    return

visited_a[0]=True
dfs(0,a[0])
visited_b[0]=True
dfs(0,b[0])
print('No')
    
