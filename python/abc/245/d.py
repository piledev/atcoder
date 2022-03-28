n,m=map(int,input().split())
a=list(map(int,input().split()))
c=list(map(int,input().split()))
b=[0]*(m+1)

for i in range(m+1):
    div=c[len(c)-i-1]//a[-1]
    b[len(b)-i-1]=str(div)
    for j in range(n+1):
        c[len(c)-i-j-1]=c[len(c)-i-j-1]-div*a[len(a)-j-1]

ans=' '.join(b)   
print(ans)
