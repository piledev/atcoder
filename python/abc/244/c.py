n=int(input())
s=[False]*(2*n+2)
a=1
while a>0:
    for j in range(1,2*n+2):
        if s[j]:
            continue
        print(j,flush=True)
        s[j]=True
        break
    a=int(input())
    s[a]=True
