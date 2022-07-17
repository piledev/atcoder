s=input()
Q=int(input())
flip=False
from collections import deque
d=deque(s)

for i in range(Q):
    q=list(map(str,input().split()))
    if q[0]=='1':
        flip = not flip
    else:
        if q[1]=='1':
            if not flip:
                d.appendleft(q[2])
            else:
                d.append(q[2])
        else:
            if flip:
                d.appendleft(q[2])
            else:
                d.append(q[2])

if flip:
    ans=list()
    for i in range(len(d)):
        ans.append(d.pop())
    print(''.join(ans))
    
print(''.join(d))

        





