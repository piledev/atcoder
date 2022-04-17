n=int(input())
a=list(map(int,input().split()))
q=int(input())

idx=dict()
for i in range(n):
    t=list()
    if a[i] in idx :
        t=idx[a[i]]
    t.append(i)
    idx[a[i]]=t

for i in range(q):
    l,r,x=map(int,input().split())
    l-=1
    r-=1
    if x in idx:
        cnt=0
        j=0
        print(f'idx[{x}]={idx[x]}')
        print(f'idx[{x}][{j}]={idx[x][j]}')
        print(idx[x][j])
        print(l)

        while idx[x][j]<l:
            cnt+=1
            j+=1

        j=len(idx[x])-1
        while r<idx[x][j]:
            cnt+=1
            j-=1
        print(len(idx[x])-cnt)
    else:
        print(0)
        


    


