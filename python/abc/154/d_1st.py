def expected_value(x):
    return((1+x)/2)

n, k = map(int,input().split())
p = list(map(int,input().split()))
# expected value: 期待値
p_ex=[]
for i in range(n):
    p_ex.append(expected_value(p[i]))

# cumulative sum: 累積和
p_ex_cum=[p_ex[0]]
for i in range(1,n):
    p_ex_cum.append(p_ex_cum[i-1]+p_ex[i])

ans=p_ex_cum[k-1]
for i in range(k,n):
    # 区間和
    interval_sum=p_ex_cum[i]-p_ex_cum[i-k]
    ans=max(ans,interval_sum)

print(ans)
