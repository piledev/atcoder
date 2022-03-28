n=int(input())

mod=10**9+7

# 全体の数
a_all=pow(10,n,mod)

# 0 を含まない数列の数
a_0=pow(9,n,mod)
# 9 を含まない数列の数
a_9=pow(9,n,mod)

# 0 と 9 の両方を含まない
a_0_9=pow(8,n,mod)

ans = a_all - (a_0 + a_9 - a_0_9)
ans%=mod
print(ans)

