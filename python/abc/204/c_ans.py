n,m=map(int,input().split())
connect=[[] for i in range(n+1)]
for i in range(m):
    a,b=map(int,input().split())
    connect[a].append(b)

from collections import deque
def bfs(start):
    visited=[False]*(n+1)
    que=deque()
    que.append(start)

    # Setting visited and count shuould be at que.append.
    # If at que.popleft, that can be duplicated setting.
    visited[start]=True
    count=1
    while len(que)>0:
        curr=que.popleft()
        for v in connect[curr]:
            if not visited[v]:
                visited[v]=True
                count+=1
                que.append(v)
    return count

ans=0
for i in range(1,n+1):
    ans+=bfs(i)

print(ans)

