h,n=map(int,input().split())
dp=[10**10]*(h+1)
for i in range(n):
    a,b=map(int,input().split())
    j=0
    while j*a<h:
        j+=1
        if j*a<h:
            dp[j*a]=min(dp[j*a],j*b)
        else:
            dp[h]=min(dp[h],j*b)
        
    print(dp)
print(dp[h])

