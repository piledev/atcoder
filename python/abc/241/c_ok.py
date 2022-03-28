yes="Yes"
no="No"
black="#"

n=int(input())
s=[]
for i in range(n):
    st=list(input())
    s.append(st)

for i in range(n):
    maxblackcount=0
    for j in range(n):
        if j+5<n:
            target=s[i][j:j+6]
            maxblackcount=max(maxblackcount,target.count(black))
        if i+5<n:
            target=[s[i][j],s[i+1][j],s[i+2][j],s[i+3][j],s[i+4][j],s[i+5][j]]
            maxblackcount=max(maxblackcount,target.count(black))
        if j+5<n and i+5<n:
            target=[s[i][j],s[i+1][j+1],s[i+2][j+2],s[i+3][j+3],s[i+4][j+4],s[i+5][j+5]]
            maxblackcount=max(maxblackcount,target.count(black))
        if 0<=j-5 and i+5<n:
            target=[s[i][j],s[i+1][j-1],s[i+2][j-2],s[i+3][j-3],s[i+4][j-4],s[i+5][j-5]]
            maxblackcount=max(maxblackcount,target.count(black))
        if maxblackcount>=4:
            print(yes)
            exit()
print(no)


            

