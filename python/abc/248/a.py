s=list(input())
m=[0]*10
for i in range(9):
    m[int(s[i])]=1

for i in range(10):
    if not m[i]:
        print(i)

