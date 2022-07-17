n=int(input())
A=list()
for i in range(n):
    a=list()
    for j in range(i+1):
        if j==0 or j==i:
            a.append(1)
        else:
            a.append(A[i-1][j-1]+A[i-1][j])
    A.append(a)
    ans=' '.join(map(lambda x:str(x),a))
    print(ans)



