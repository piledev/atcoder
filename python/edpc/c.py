n=int(input())
happiness=list()
for i in range(n):
    a,b,c=map(int,input().split())
    happiness.append([a,b,c])

dp=[[0]*3 for i in range(n)]
dp[0]=happiness[0]

for i in range(1,n):
    for j in range(3):
        for k in range(3): 
            if j==k:
                continue
            dp[i][j]=max(dp[i][j],dp[i-1][k]+happiness[i][j])

print(max(dp[-1]))


