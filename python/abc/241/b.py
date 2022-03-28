n,m=map(int,input().split())
a=list(map(int,input().split()))
b=list(map(int,input().split()))
used=[False]*n
a.sort()
b.sort()

stop=0
for i in range(m):
    ok=False
    for j in range(stop,n):
        if (not used[j]) and b[i]==a[j]:
            ok=True
            used[j]=True
            break
    if not ok:
        print("No")
        exit()
print("Yes")

        



