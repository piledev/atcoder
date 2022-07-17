n=int(input())
d={y:{x:0 for x in range(10)} for y in range(10)}
for i in range(n):
    s=list(map(lambda x:int(x),list(input())))
    for i in range(10):
        d[s[i]][i]+=1

ans=10**15
# i を揃える場合
for i in range(10):
    idx_cnt_max=-1
    idx_cnt_max_idx=-1
    for j in range(10):
        if idx_cnt_max<=d[i][j]:
            idx_cnt_max=d[i][j]
            if idx_cnt_max_idx<=j:
                idx_cnt_max_idx=j
    ans=min(ans,idx_cnt_max_idx+10*(idx_cnt_max-1))
print(ans)

