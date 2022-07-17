x,a,d,n=map(int,input().split())

left=min(a,a+d*(n-1))
right=max(a,a+d*(n-1))
ans=0
if x <= left:
    ans=abs(left-x)
elif right <= x:
    ans=abs(x-right)
else:
    mod=abs(x-left)%abs(d)
    ans=min(abs(mod),abs(abs(d)-abs(mod)))
print(ans)


