n=int(input())
mx=200001
ll=[0]*mx
rl=[0]*mx

for i in range(n):
    l,r=map(int,input().split())
    ll[l]+=1
    rl[r]+=1


l=0
r=0
now=0
for i in range(len(ll)):
    if now == 0 and ll[i]>0:
        l=i
    if now > 0 and now + ll[i]- rl[i]==0:
        r=i
        print(l,r)
    now += ll[i]
    now -= rl[i]


