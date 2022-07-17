n=int(input())
a=list(map(int,input().split()))

ans=0
for i in range(n):
    total=sum(a[i:n])
    if (total)>3:
        ans+=1

print(ans)
    
    
