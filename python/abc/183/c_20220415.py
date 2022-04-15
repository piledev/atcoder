n,k=map(int,input().split())
t=list()
for i in range(n):
    r=list(map(int,input().split()))
    t.append(r)

ans=0

def rec(total,visited):
    global ans
    if len(visited)==n:
        if total+t[visited[-1]][0]==k:
            ans+=1

    for i in range(n):
        if i in visited:
            continue
        else:
            visited.append(i)
            rec(total+t[visited[-2]][i],visited)
            visited.pop()

visited=[0]
rec(0,visited)
print(ans)

