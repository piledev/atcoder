n = int(input())
if n%3==0:
    print(0)
    exit()

dic = {0:0,1:0,2:0}
ns = str(n)
for x in ns:
   dic[int(x)%3] += 1  

ans = -1
if n%3==1:
    if dic[1] > 0:
        ans = 1
    elif dic[2] > 1:
        ans = 2 
else:
    if dic[2] > 0:
        ans = 1
    elif dic[1] > 1:
        ans = 2

if ans == len(ns):
    ans = -1

print(ans)

