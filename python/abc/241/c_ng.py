n=int(input())
s=[]
for i in range(n):
    st=list(input())
    s.append(st)

for i in range(n):
    for j in range(n):
        yoko=[]
        tate=[]
        naname1=[]
        naname2=[]
        for k in range(6):
            if i+k < n:
                tate.append(s[i+k][j])
            if j+k < n:
                yoko.append(s[i][j+k])
            if i+k < n and j+k < n:
                naname1.append(s[i+k][j+k])
            if 0<=i-5 and j+k < n:  
                naname2.append(s[i-k][j+k])
        if tate.count("#")>=4 or yoko.count("#")>=4 or naname1.count("#")>=4 or naname2.count("#")>=4:
            print("Yes")
            exit()
print("No")

            

