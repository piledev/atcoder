import math
n=int(input())

sosu=list()
# for i in range(math.ceil((n**(-3)))):
for i in range(2,n):
    j=2
    while j**2<=i:
        if i%j==0:
            break
        j+=1
    else:
        sosu.append(i)
print(sosu)
