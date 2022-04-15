import math
n=int(input())
li=[]
for a in range(2,int(math.sqrt(n))+1):
    b=2
    while a**b<=n:
        li.append(a**b)
        b+=1
        
s=set(li)
print(n-len(s))

