n,W=map(int,input().split())
inf=100**15
V=1000*n # 合計価値の最大値
dp=list()
dp=[[0]+[inf]*100000 for i in range(n+1)]

for i in range(1,n+1):
    w,v=map(int,input().split())
    for j in range(1,V+1):
        if dp[i-1][j]<inf:
            dp[i][j]=min(dp[i][j],dp[i-1][j])
        if v<=j and dp[i-1][j-v]<inf:
            dp[i][j]=min(dp[i][j],dp[i-1][j-v]+w)
            # print(f'dp[{i}][{j}]={dp[i][j]}')

for i in range(V,0,-1):
    if dp[-1][i]<=W:
        print(i)
        exit()

