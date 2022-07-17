import heapq
Q=int(input())
ha=list()
heapq.heapify(ha)
hd=list()
heapq.heapify(hd)
d=dict()

for i in range(Q):
    q=list(map(int,input().split()))
    if q[0]==1:
        x=q[1]
        if x not in d:
            d[x]=1
            heapq.heappush(ha, x)
            heapq.heappush(hd, -1*x)
        else:
            d[x]+=1

    elif q[0]==2:
        x=q[1]
        c=q[2]
        if x in d:
            d[x]-=min(d[x],c)
            if d[x]==0:
                del(d[x])

    else:
        mx=-1*heapq.heappop(hd)
        while mx not in d:
            mx=-1*heapq.heappop(hd)
        heapq.heappush(hd, -1*mx)

        mn=heapq.heappop(ha)
        while mn not in d:
            mn=heapq.heappop(ha)
        heapq.heappush(ha, mn)
        
        print(mx-mn)




