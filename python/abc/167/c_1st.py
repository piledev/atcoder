n,m,x=map(int,input().split())
c=list()
a=list()
for i in range(n):
    s=list(map(int,input().split()))
    c.append(s[0])
    a.append(s[1:])

# 最小を求めるので
ans=10**15


# bit全探索
"""
<<: シフト演算。1bit左にずらす。
n=4の例） 
    1<<4 = 10000b = 1111b + 1b
    つまり range(1<<4) は 0000b ~ 1111b をすべて試すという意味になる
"""
for i in range(1<<n):
    cost=0
    skill=[0]*m

    """
    i を右にずらしながら 1 桁目が 1 かどうかを確認する。
    1 なら「買う」なので、コストと理解度を加算する。
    初回の shift は0なので実質ずらしてない状態。
    最後の shift は N-1 なので最後の1つを確認することになる。
    n=4の例）
        1000b を 3 回シフトすると 4 個目の 1 が一桁目に来る
    """
    for shift in range(n):
        if i>>shift&1==1:
            cost+=c[shift]
            for j in range(m):
                skill[j]+=a[shift][j]
        if min(skill)>=x:
            ans=min(ans,cost)

if ans==10**15:
    print(-1)
else:
    print(ans)
