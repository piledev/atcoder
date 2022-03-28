n,m=map(int,input().split())
connect=[[] for i in range(n)]
for i in range(m):
    a,b=map(int,input().split())
    a-=1
    b-=1
    connect[a].append(b)

from collections import deque

def bfs(start):
    visited=[False]*n
    visited[start]=True
    cnt=1
    que=deque()
    que.append(start)
    while len(que)>0:
        current_city=que.popleft()
        for to_city in connect[current_city]:
            if not visited[to_city]:
                visited[to_city]=True
                cnt+=1
                que.append(to_city)
    return cnt
        
ans=0
for i in range(n):
    ans+=bfs(i)

print(ans)


