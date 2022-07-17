n=int(input())
a=list(map(int,input().split()))

all=n*(n-1)*(n-2)//6
double=0
triple=0

cnts=dict()
for v in a:
    if not v in cnts:
        cnts[v]=1
    else:
        cnts[v]=cnts[v]+1

for k,v in cnts.items():
    print(k,v)
    if v<=1:
        continue
    if v==2:
        if a[-1]==v:
            double+=2
        else:
            double+=3
    elif v==3:
        triple+=1
        if a[-1]==v:
            double+=2
        else:
            double+=3
    else:
        triple+=(v*(v-1)*(v-2))//6
        if a[-1]==v:
            double+=(v*(v-1))//2

        else:
            double+=(v*(v-1))//2

print(all,double,triple)
ans=all-double-triple
print(ans)

