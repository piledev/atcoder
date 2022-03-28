s=input().split()
t=input().split()

cnt=0
for i in range(3):
    if s[i]==t[i]:
        cnt+=1
if cnt==1:
    print('No')
    exit()

print('Yes')

