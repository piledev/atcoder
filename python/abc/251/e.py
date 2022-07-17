n=int(input())
a=list(map(int,input().split()))

dp1=[0 for i in range(n)]
dp2=[0 for i in range(n)]

# 最初の行動を行わない場合
dp1[1]=a[1]
dp1[2]=a[1]
for i in range(2,n):
    if i==n-1:
        dp1[i]=dp1[i-1]+a[i]
        break

    if a[i-1]<a[i]:
        dp1[i]=dp1[i-2]+a[i-1]
    else:
        dp1[i]=dp1[i-1]+a[i]
        dp1[(i+1)%n]=dp1[i]

dp2[0]=a[0]
dp2[1]=a[0]
for i in range(2,n):
    if a[i-1]<a[i]:
        dp2[i]=dp2[i-2]+a[i-1]
    else:
        dp2[i]=dp2[i-1]+a[i]
        dp2[(i+1)%n]=dp2[i]


print(dp1)
print(dp2)

# 自分の回答
# if a[n-1]<a[0]:
#     lastanimal=1
#     dp[0]=a[n-1]
# else:
#     dp[0]=a[0]
#     dp[1]=a[0]

# for i in range(1,n-lastanimal):
#     if dp[i]<10**15:
#         continue
#     else:
#         if a[i-1]<a[i]:
#             dp[i]=dp[i-1]+a[i-1]
#         else:
#             dp[i]=dp[i-1]+a[i]
#             dp[i+1]=dp[i]

# print(dp[n-1-lastanimal])
# print(dp)

