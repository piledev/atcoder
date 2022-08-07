n,l,r=map(int,input().split())
a=list(map(int,input().split()))

x=0
y=0

lsum=0
lmax=0
lmaxidx=0
for i in range(n):
    lsum+=a[i]
    if lmax<lsum-l*(i+1):
        lmax=lsum-l*(i+1)
        lmaxidx=i
if lmax>0:
    x=lmaxidx
    for i in range(x+1):
        a[i]=l

rsum=0
rmax=0
rmaxidx=0
for i in range(n):
    rsum+=a[n-1-i]
    if rmax<rsum-r*(i+1):
        rmax=rsum-r*(i+1)
        rmaxidx=i
if rmax>0:
    y=rmaxidx
    for i in range(y+1):
        a[n-1-i]=r

# print('x:',x,'y:',y)

ans=sum(a)
print(ans)
