x,k,d=map(int,input().split())
x=abs(x)
ans=10**15
cur=x
if d*k<=x:
    ans=x-(d*k)
else:
    k-=x//d
    curr=x-d*(x//d)
    if k%2==0:
        ans=curr
    else:
        ans=d-curr

print(ans)




