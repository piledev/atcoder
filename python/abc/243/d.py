n,x=map(int,input().split())
s=input()

t=list()
for c in s:
    if c=='U' and len(t)>0 and (t[-1]=='L' or t[-1]=='R'):
        t.pop()
    else:
        t.append(c)


for c in t:
    if c=='U':
        x=x//2
        continue
    x=x*2
    if c=='R':
        x+=1

print(x)

