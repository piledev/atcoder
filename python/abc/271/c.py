n=int(input())
a=list(map(int,input().split()))
a.sort()

aset = set(a)
dupcnt = len(a) - len(aset)

ans=0
for i in range(1,n+1):
    if i in aset and a[-1]>=i:
        ans=i
    else:
        if dupcnt > 1:
            dupcnt-=2
            ans=i
        elif dupcnt == 1:
            if a[-1]>i:
                dupcnt = 0
                a.pop()
                ans=i
            else:
                break
        elif len(a)>1 and a[-2]>i:
            a.pop()
            a.pop()
            ans=i
        else:
            break

print(ans)

