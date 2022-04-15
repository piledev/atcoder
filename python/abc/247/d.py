from collections import deque
li=deque()
Q=int(input())
for i in range(Q):
    s=list(map(int,input().split()))
    if s[0]==1:
        li.append([s[1],s[2]])
    else:
        c2=s[1]
        sum=0
        while c2>0 and len(li)>0:
            x,c1=map(int,li.popleft())
            if c1<=c2:
                sum+=x*c1
                c2-=c1
            else:
                sum+=x*c2
                c1-=c2
                li.appendleft([x,c1])
                c2=0
        print(sum)

