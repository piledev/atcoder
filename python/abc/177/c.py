n=int(input())
a=list(map(int,input().split()))
mod=1000000007
ans=0
total=sum(a)
for i in range(n-1):
    total-=a[i]
    ans+=a[i]*total
    ans%=mod

print(ans)
