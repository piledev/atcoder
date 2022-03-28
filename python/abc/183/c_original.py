N,K = map(int,input().split())
ANS = 0
T = []
arrived=[False for i in range(N)]
for i in range(N):
    # T.append(list(map(int,input().split())))
    T.append(map(int,input().split()))

def rec(arrived, depth, total, current):
    global ANS
    if depth == N:
        total += T[current][0]
        if total == K:
            ANS += 1
        return

    for i in range(len(arrived)) :
        if arrived[i]:
            continue
        else:
            arrived[i]=True
            rec(arrived,depth+1,total+int(T[current][i]),i)
            arrived[i]=False
    return

arrived[0]=True
rec(arrived, 1, 0, 0)    

print(ANS)
