n,x,y=map(int,input().split())
a=list(map(int,input().split()))

posx,posy,b=-1,-1,-1
ans=0

for i in range(n):
    if a[i]==x:
        posx=i
    if a[i]==y:
        posy=i
    if a[i]<y or x<a[i]:
        b=i

    if b<min(posx,posy):
        ans+=min(posx,posy)-b
print(ans)

