n = int(input())
m = set()

a = 2 
while a**2 <= n:
    b = 2
    while a**b <= n:
        m.add(a**b)
        b += 1
    a += 1
ans = n - len(m)
print(ans)

