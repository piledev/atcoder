def divisor_list(n):
    divisor = []
    i=1
    while i**2 <= n:
        if n%i==0:
            divisor.append(i)
            if i != n//i:
                divisor.append(n//i)
        i += 1
    divisor.sort()
    return divisor

n = int(input())
ans = divisor_list(n)

for v in ans:
    print(v)
