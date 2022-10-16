n=int(input())
a=list()
for i in range(n):
    a.append(input())

s=set()
# horizontal check
for i in range(n):
    for j in range(n):
        h=a[i][j:]+a[i][:j]
        s.add(h)

# vertical
print(s)

    




