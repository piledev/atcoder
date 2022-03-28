a,b,c,x=map(int, input().split())

ans = float(0)
if x <= a:
    ans = float(1)
elif x <= b:
    ans = c/(b-a)

print(ans)
    

