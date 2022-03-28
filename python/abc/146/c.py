a,b,x=map(int,input().split())

def d(n):
    return len(str(n))
    
def valueof(n):
    return a*n+b*d(n)

ans=0
l=0
r=10**18
count=0
while l+1<r:
    n=(l+r)//2
    value=valueof(n)
    if value==x:
        print(n)
        exit()
    elif value<x:
        ans=max(ans,n)
        l=n
    else:
        r=n
if 10**9<ans:
    ans = 10**9
print(ans)

    
