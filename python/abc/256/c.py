h1,h2,h3,w1,w2,w3=map(int,input().split())
h=[h1,h2,h3]
w=[w1,w2,w3]

ans=0
for i0 in range(1,h[0]-1):
    cells=[[0]*3 for _ in range(3)]
    cells[0][0]=i0
    for j0 in range(1,h[0]-i0):
        cells[0][1]=j0
        cells[0][2]=h[0]-i0-j0

        for i1 in range(1,h[1]-1):
            cells[1][0]=i1
            for j1 in range(1,h[1]-i1):
                cells[1][1]=j1
                cells[1][2]=h[1]-i1-j1
                for x in range(3):
                    cells[2][x]=w[x]-cells[0][x]-cells[1][x]

                # print(cells, sum(cells[2]),h[2])
                if sum(cells[2])==h[2] :
                    if cells[2][0]>0 and cells[2][1]>0 and cells[2][2]>0:
                        ans+=1
print(ans)



