n=int(input())
p=list(map(int,input().split()))

parent=n
for i in range(len(p)):
    parent-=2
    parent=p[parent]
    if parent==1:
        print(i+1)
        exit()
    


