n=int(input())
x=list()
y=list()
for i in range(n):
    x1,y1=map(int,input().split())
    x.append(x1)
    y.append(y1)

for i in range(n-2):
    for j in range(i+1,n-1):
        for k in range(j+1,n):
            if (y[k]-y[i])*(x[j]-x[i])==(y[j]-y[i])*(x[k]-x[i]):
                print('Yes')
                exit()
print('No')
