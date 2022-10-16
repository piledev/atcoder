n,q=map(int,input().split())

aa = list()
for i in range(n):
    a = list(map(int,input().split()))[1:]
    aa.append(a)

for i in range(q):
    s,t=map(int,input().split())
    print(aa[s-1][t-1])


