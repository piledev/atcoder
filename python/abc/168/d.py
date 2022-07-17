n,m=map(int,input().split())
conn=[[] for i in range(n+1)]
for i in range(m):
    a,b=map(int,input().split())
    conn[a].append(b)
    conn[b].append(a)

visited=[False]*(n+1)
signs=[0]*(n+1)

from collections import deque
que=deque()

def bfs(start):
    que.append(start)
    visited[start]=True
    while len(que)>0:
        curr=que.popleft()
        for v in conn[curr]:
            if not visited[v]:
                que.append(v)
                visited[v]=True
                signs[v]=curr

bfs(1)

print('Yes')
for i in range(2,n+1):
    print(signs[i])




