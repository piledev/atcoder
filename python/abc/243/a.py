v,a,b,c=map(int,input().split())

t=v
while 1>0:
    t-=a
    if t<0:
        print('F')
        exit()
    t-=b
    if t<0:
        print('M')
        exit()
    t-=c
    if t<0:
        print('T')
        exit()

