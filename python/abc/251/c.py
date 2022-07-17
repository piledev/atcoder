n=int(input())

poems=dict() # key: poem, value: original's (index, point)

for i in range(n):
    s,t=map(str,input().split())
    t=int(t)
    if s in poems.keys():
        continue
    poems[s]=(i,t)

max_point=-10**15
max_point_idx=10**15
for v in poems.values():
    idx, point = v
    if point == max_point:
        if idx < max_point_idx:
            max_point_idx = idx
    elif point > max_point:
        max_point = point
        max_point_idx = idx

print(max_point_idx+1)




