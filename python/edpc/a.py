n=int(input())
h=list(map(int,input().split()))
dp=[10**15]*n
dp[0]=0
dp[1]=abs(h[1]-h[0])
for i in range(2,n):
    dp[i]=min(dp[i],dp[i-2]+abs(h[i]-h[i-2]))
    dp[i]=min(dp[i],dp[i-1]+abs(h[i]-h[i-1]))

print(dp[-1])



