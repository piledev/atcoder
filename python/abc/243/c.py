n=int(input())
f=list()
for i in range(n):
    f1=list(map(int,input().split()))
    f.append(f1)
s=input()

for i in range(n):
    f[i].append(s[i])

f.sort(key=lambda x:x[1])

x=0
y=1
lr=2
minmax=list()
rmin=10**15
lmax=-10**15
for i in range(n):
    if i==0:
        minmax.append([rmin,lmax])
        if f[i][lr]=='L':
            minmax[i][1]=f[i][x]
        else:
            minmax[i][0]=f[i][x]
        continue
    elif f[i][y]!=f[i-1][y]:
        minmax.append([rmin,lmax])

    if f[i][lr]=='L':
        minmax[len(minmax)-1][1]=max(minmax[len(minmax)-1][1],f[i][x])
    else:
        minmax[len(minmax)-1][0]=min(minmax[len(minmax)-1][0],f[i][x])

for v in minmax:
    if v[0] < v[1]:
        print("Yes")
        exit()

print("No")

