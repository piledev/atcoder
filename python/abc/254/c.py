n,k=map(int,input().split())
a=list(map(int,input().split()))

if k == 1:
    print('Yes')
    exit()

b=[[] for i in range(k)]
for i in range(n):
    b[i%k].append(a[i])

for i in range(len(b)):
    b[i].sort()

for i in range(n):
    a[i]=b[i%k][i//k]
    if i>0 and a[i-1]>a[i]:
        print('No')
        exit()
print('Yes')

    
