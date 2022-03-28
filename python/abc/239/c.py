x1,y1,x2,y2=map(int,input().split())

def koshiten(x,y):
    li=list()
    li.append((x-2,y+1))
    li.append((x-1,y+2))
    li.append((x+1,y+2))
    li.append((x+2,y+1))
    li.append((x-2,y-1))
    li.append((x-1,y-2))
    li.append((x+1,y-2))
    li.append((x+2,y-1))
    return li


l1 = koshiten(x1,y1)
l2 = koshiten(x2,y2)

ans="No"
count=0
for x in l1:
    if x in l2:
        ans="Yes"

print(ans)
