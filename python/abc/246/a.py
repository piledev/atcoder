x=list()
y=list()
for i in range(3):
    a,b=map(int,input().split())
    x.append(a)
    y.append(b)

x.sort()
y.sort()

X,Y=0,0
if x[0]==x[1]:
    X=x[2]
else:
    X=x[0]
        
if y[0]==y[1]:
    Y=y[2]
else:
    Y=y[0]

print(X,Y)


