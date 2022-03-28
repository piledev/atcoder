n=int(input())
t=input()
b=1
current=[0,0]
for i in range(n):
    if t[i]=='S':
        if b==1:
            current[0]+=1
        elif b==2:
            current[1]-=1
        elif b==3:
            current[0]-=1
        else:
            current[1]+=1
    else:
        b=(b+1)%4
print(current[0],current[1])
