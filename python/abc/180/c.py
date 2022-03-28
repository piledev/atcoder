import math
n = int(input())

a = []
for i in range(int(math.sqrt(n))+1):
    if n % (i+1) == 0 and (i+1)<=n//(i+1):
        a.append(i+1)
        if n//(i+1) not in a:
            a.append(n//(i+1)) 
a.sort()
for i in range(len(a)):
    print(a[i]) 

