r,c = map(int,input().split())
b='black'
w='white'
ans=b
if r==2 or r==14:
    if 2<=c<=14:
        ans=w
elif r==4 or r==12:
    if c%2==0  or 4<=c<=12:
        ans=w
elif r==6 or r==10:
    if c%2==0  or 6<=c<=10:
        ans=w
elif r==8:
    if c%2==0:
        ans=w
elif r==3 or r==13:
    if c==2 or c==14:
        ans=w
elif r==5 or r==11:
    if c==2 or c==4 or c==12 or c==14:
        ans=w
elif r==7 or r==9:
    if c==2 or c==4 or c==6 or c== 10 or c==12 or c==14:
        ans=w

print(ans)

