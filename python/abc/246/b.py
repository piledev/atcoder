import math
a,b=map(int,input().split())
x=math.sqrt(1/(a**2+b**2))

print(a*x,b*x)
