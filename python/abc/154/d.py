n,k=map(int,input().split())
p=list(map(int,input().split()))
hopes=list()
for i in range(n):
    hopes.append(((p[i]+1)/2))

totals=list()
total=sum(hopes[0:k])
totals.append(total)

for i in range(n-k):
    total=totals[i]-hopes[i]+hopes[i+k]
    totals.append(total)

print(max(totals))


