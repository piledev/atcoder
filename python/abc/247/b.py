n=int(input())
s=list()
t=list()
for i in range(n):
    x,y=map(str,input().split())
    s.append(x)
    t.append(y)

ans='Yes'
for i in range(n):
    sb=True
    tb=True
    for j in range(n):
        if i==j:
            continue
        if s[i]==s[j] or s[i]==t[j]:
            sb=False
        if t[i]==s[j] or t[i]==t[j]:
            tb=False
        if sb==False and tb==False:
            ans='No'
            print(ans)
            exit()
print(ans)
