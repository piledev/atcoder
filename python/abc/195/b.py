a,b,w=map(int,input().split())
cannot='UNSATISFIABLE'
w*=1000

fr=w//b
to=w//a
if w%a>0:
    to+=1

mi=10**15
ma=-1*(10**15)
for i in range(fr,to+1):
    if a*i<=w<=b*i:
        mi=min(mi,i)
        ma=max(ma,i)
if mi==10**15:
    print(cannot)
else:
    print(mi,ma)
    

