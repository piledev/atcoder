n,m=map(int,input().split())

def rec(s,v):
    s2=s.copy()
    s2.append(v)
    if len(s2) == n:
        # print(' '.join(map(str,s2)))
        print(*s2)
        return
    for i in range(v+1,m+1):
        # print(f'rec({s2},{i})')
        rec(s2,i)

s=list()
for i in range(1,m-n+2):
    rec(s,i)


