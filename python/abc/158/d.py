from collections import deque

inv=False
s=deque(list(input()))
q=int(input())
for i in range(q):
    query=list(map(str,input().split()))
    if query[0]=="1":
        inv = not inv
    else:
        if not inv:
            if query[1]=="1":
                s.appendleft(query[2])
            else:
                s.append(query[2])
        else:
            if query[1]=="1":
                s.append(query[2])
            else:
                s.appendleft(query[2])

if inv:
    s.reverse()

print("".join(s))


