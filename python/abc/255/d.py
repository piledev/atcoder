n,q=map(int,input().split())
a=list(map(int,input().split()))
a.sort()

def bs(x):
    l=0
    r=len(a)-1
    if a[l] == x or x == a[r]:
        return 0
    elif x < a[l]:
        return a[l]-x
    elif a[r] < x:
        return x - a[r]

    while a[l] <= x <= a[r]:
        if (r-l)//2 <= x:
            l=(r-l)//2
        else:
            r=(r-l)//2
        # print(f'a[{l}]:{a[l]}, a[{r}]:{a[r]}')
        if a[l] == x or x == a[r]:
            return 0
        if l-r==1:
            return min(x-a[l],a[r]-x)

for i in range(q):
    x=int(input())
    print(bs(x))


