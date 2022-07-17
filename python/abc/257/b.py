n,k,q=map(int,input().split())
a=list(map(int,input().split()))
l=list(map(int,input().split()))

for i in range(len(l)):
    a.sort()
    v=l[i]
    if a[v-1]==n:
        continue
    else:
        if (a[v-1]+1) in a:
            continue
        else:
            a[v-1]+=1

b=list(map(lambda x:str(x),a))
print(' '.join(b))

