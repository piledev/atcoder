n=int(input())
s=list(input())
q=int(input())
flip=False
for i in range(q):
    t,a,b=map(int,input().split())
    a-=1
    b-=1
    if t==1:
        if not flip:
            s[a],s[b]=s[b],s[a]
        else:
            if a<n:
                a+=n
            else:
                a-=n
            if b<n:
                b+=n
            else:
                b-=n
            s[a],s[b]=s[b],s[a]
    elif t==2:
        flip=not flip
if flip:
    s=s[n:]+s[:n]
print(''.join(s))

    
