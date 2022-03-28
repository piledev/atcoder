taka="Takahashi"
aoki="Aoki"
sosu=[True]*201

a,b,c,d=map(int,input().split())

for i in range(201):
    j=2
    while j**2<=i:
        if i%j==0:
            sosu[i]=False
            break
        j+=1

ans=aoki
for i in range(a,b+1):
    not_sosu=True
    for j in range(c,d+1):
        if sosu[i+j]==True:
            not_sosu=False
            break
    if not_sosu:
        ans=taka

print(ans)









