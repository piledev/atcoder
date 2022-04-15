n=int(input())
m=int(n**(1/3))
ans=10**30
for a in range (m+1):
    for b in range (m,-1,-1):
        x=a**3+(a**2)*b+a*(b**2)+b**3
        # print(a,b,x)
        if x>=n:
            ans=min(ans,x)
print(ans)


