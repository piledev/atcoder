n,W=map(int,input().split())
dp=[[0]*(W+1) for i in range(n+1)]
for i in range(1,n+1):
    w,v=map(int,input().split())
    for j in range(1,W+1):
        if w<=j:
            dp[i][j]=max(dp[i-1][j],dp[i-1][j-w]+v)
        else:
            dp[i][j]=dp[i-1][j]

print(dp[-1][-1])

