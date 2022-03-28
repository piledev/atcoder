n,k=map(int,input().split())
a=list(map(int,input().split()))

current_town=1
visited=[current_town]
loop=0
for i in range(0,k):
    current_town = a[current_town-1]
    visited.append(current_town)
    if i>0 and current_town==1:
        loop=i+1
        break
ans=visited[k%loop]
print(ans)

