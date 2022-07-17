n,k=map(int,input().split())
r,s,p=map(int,input().split())
t=list(input())

points=[0]*n
ans=0
for i in range(n):
    if i<k:
        if t[i]=='r':
            points[i]=p
        elif t[i]=='s':
            points[i]=r
        else:
            points[i]=s
    else:
        if t[i]==t[i-k] and points[i-k]>0:
            pass
        else:
            if t[i]=='r':
                points[i]=p
            elif t[i]=='s':
                points[i]=r
            else:
                points[i]=s

print(sum(points))




