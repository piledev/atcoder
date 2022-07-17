h,w,k=map(int,input().split())
c=[]
for i in range(h):
    c.append(list(input()))

ans=0
for row_red in range(1<<h):
    for col_red in range(1<<w):
        black=0
        for row in range(h):
            for col in range(w):
                if (row_red>>row)&1==0 and (col_red>>col)&1==0 and c[row][col]=="#":
                    black+=1
        if black==k:
            ans+=1
print(ans)

