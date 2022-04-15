n,k,x=map(int,input().split())
a=list(map(int,input().split()))

a.sort(reverse=True)
for i in range(n):
    div=a[i]//x
    if div>0:
        usable=min(a[i]//x,k)
        k-=usable
        a[i]-=usable*x
        if k==0:
            print(sum(a))
            exit()
    else:
        break

a.sort(reverse=True)
for i in range(n):
    if k>0:
        k-=1
        a[i]=0
    else:
        break
print(sum(a))

