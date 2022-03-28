n = int(input())
p = []
for i in range(n):
    x, y = map(int,input().split())
    p.append((x,y))

for i in range(n-2):
    x1, y1 = p[i]
    for j in range(i+1,n-1):
        x2, y2 = p[j]
        for k in range(j+1,n):
            x3, y3 = p[k]
            if (y3-y1)*(x2-x1)==(y2-y1)*(x3-x1):
                print("Yes")
                exit()
print("No")



