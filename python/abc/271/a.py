n = int(input())

def change(i):
    if i < 10:
        return str(i)
    elif i == 10:
        return 'A'
    elif i == 11:
        return 'B'
    elif i == 12:
        return 'C'
    elif i == 13:
        return 'D'
    elif i == 14:
        return 'E'
    else:
        return 'F'

ten = n // 16
one = n % 16

ten = change(ten)
one = change(one)

ans = ten + one
print(ans)


