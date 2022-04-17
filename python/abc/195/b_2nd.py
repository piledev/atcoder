a,b,w=map(int,input().split())
w*=1000
mi=10**15
ma=-10**15

for i in range(1,w//a+1):
    if a*i<=w<=b*i:
        mi=min(mi,i)
        ma=max(ma,i)
if mi==10**15:
    print("UNSATISFIABLE")
    exit()
print(mi,ma)



