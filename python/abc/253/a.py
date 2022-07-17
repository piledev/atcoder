a,b,c=map(int,input().split())

if a<=b<=c:
    print('Yes')
elif c<=b<=a:
    print('Yes')
else:
    print('No')
