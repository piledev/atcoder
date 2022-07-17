n=int(input())
s=input()
w=list(map(int,input().split()))

ans=0
people=list()
for i in range(n):
    people.append((w[i],s[i]))
    if s[i]=='1':
        ans+=1

people.sort()

x=ans
for i in range(n):
    if people[i][1]=='1':
        x-=1
    else:
        x+=1
    if i<n-1:
        if people[i][0] != people[i+1][0]:
            ans=max(ans,x)
    else:
        ans=max(ans,x)

print(ans)
