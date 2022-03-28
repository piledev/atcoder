n=int(input())
a=list(map(int,input().split()))
b=list(map(int,input().split()))

x=0
y=0
for i in range(n):
    if a[i]==b[i]:
        x+=1
    elif a[i] in b:
        y+=1

print(x)
print(y)


