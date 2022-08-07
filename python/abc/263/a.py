c=input().split()

cset = set(c)
if len(cset)!=2:
    print('No')
    exit()

a=c[0]
count=0
for v in c:
    if a==v:
        count+=1

if 2 <= count <=3:
    print('Yes')
else:
    print('No')
