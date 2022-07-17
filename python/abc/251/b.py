n,w=map(int,input().split())
a=list(map(int,input().split()))

ans=list()
for i in range(n):
    if a[i]<=w:
        ans.append(a[i]) 
    for j in range(i+1,n):
        if a[j]<=w:
            ans.append(a[j])
        if a[i]+a[j]<=w:
            ans.append(a[i]+a[j])
        for k in range(j+1,n):
            if a[k]<=w:
                ans.append(a[k])
            if a[j]+a[k]<=w:
                ans.append(a[j]+a[k])
            if a[i]+a[j]+a[k]<=w:
                ans.append(a[i]+a[j]+a[k])

ans=set(ans)
print(len(ans))
            



