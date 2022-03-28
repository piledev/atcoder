def g1(x):
    ret = 0
    for i in range(0,len(x)):
        ret += int(x[i])*(10**i)
    return ret

def g2(x):
    ret = 0
    for i in range(0,len(x)):
        ret += int(x[len(x)-i-1])*(10**i)
    return ret

def f(i):
    x = list(str(i))
    x.sort()
    return g1(x) - g2(x)

n, k = map(int,input().split())
ans = n
for i in range(0,k):
    ans = f(ans)

print(ans)


