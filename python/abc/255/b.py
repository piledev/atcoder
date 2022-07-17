import math
n,k=map(int,input().split())
a=list(map(int,input().split()))

ans=-10**15
dark=list()
light=list()

for i in range(n):
    x,y=map(int,input().split())
    if i+1 in a:
        light.append([x,y])
    else:
        dark.append([x,y]) 

for d in dark:
    mini=10**15
    for l in light:
        dist=math.sqrt((d[0]-l[0])**2+(d[1]-l[1])**2)
        mini=min(mini,dist)
    ans=max(ans,mini)

print(ans)

        

    



