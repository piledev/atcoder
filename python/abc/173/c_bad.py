import copy
h,w,k=map(int,input().split())
c=list()
for i in range(h):
    c.append(list(input()))

# def myprint(t):
#     print('='*25)
#     for v in t:
#         print(v)
#     print('='*25)

ans=0
for i in range(1<<h+w):
    t=copy.deepcopy(c)
    # print('i: ',i)
    for shift in range(h+w):
        if i>>shift&1==1:
            # print('shift: ',shift)
            if shift<h:
                t[shift]=['x']*w
            else:
                for j in range(h):
                    t[j][shift-h]='x'
    count=0
    for a in t:
        for b in a:
            if b=='#':
                count+=1
    if count==k:
        # myprint(t)
        ans+=1
print(ans)
