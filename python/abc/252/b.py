n,k=map(int,input().split())
a=list(map(int,input().split()))
b=list(map(int,input().split()))

deli=0
for v in a:
    if deli<v:
        deli=v

delis=[]
for i in range(n):
    if a[i]==deli:
        delis.append(i+1)
        
for v in b:
    if v in delis:
        print('Yes')
        exit()

print('No')

    

