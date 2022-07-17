n,m,x,t,d=map(int,input().split())
f = t - (x*d)
ans = f + min(m,x)*d
print(ans)

