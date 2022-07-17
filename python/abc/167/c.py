n,m,x=map(int,input().split())
c=list()
a=list()
for i in range(n):
    t=list(map(int,input().split()))
    c.append(t[0])
    a.append(t[1:])

min_cost=10**15
for i in range(1<<n):
    cost=0
    powers=[0]*m
    for j in range(n):
        if i>>j&1==1:
            cost+=c[j]
            for k in range(m):
                powers[k]+=a[j][k]
    if min(powers)>=x:
        min_cost=min(min_cost,cost)

if min_cost==10**15:
    print(-1)
    exit()

print(min_cost)



            






