n = int(input())
s = list(input())
q = int(input())

flip = False
for i in range(q):
    t,a,b = map(int,input().split())
    if t == 1:
        if flip:
            a = (a+n)%(2*n)
            b = (b+n)%(2*n)
        s[a-1],s[b-1] = s[b-1],s[a-1]
    else:
        flip = not flip

if flip:
    s = s[n:] + s[:n]

s = "".join(s)
print(s)



