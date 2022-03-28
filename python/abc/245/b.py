n=int(input())
a=map(int,input().split())

l=[False]*2001

for v in a:
    l[v]=True

ans=2001
for i in range(n+1):
    if l[i]==False:
        ans=i
        break

print(ans)

