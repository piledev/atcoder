n=int(input())

factors=[]
i=2
temp=n
while i**2<=n:
    if temp%i==0:
        factors.append(i)
        temp=temp//i
    else:
        i+=1
count=0
for i in range(len(factors)):
    if i == 0:



