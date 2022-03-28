def judge10(x):
    x10 = str(x)
    if "7" in x10:
        return False
    return True

def judge8(x):
    while x >= 7 :
        # print(x)
        mod = x % 8
        if mod == 7:
            return False
        x = x // 8 
    return True

n = int(input())
ans = 0

for i in range(1,n+1):
    if judge10(i) and judge8(i):
        ans += 1
        
print(ans)

"""
83 を8進数に変換する場合

83 % 8 = 3
ans = "3"
83 // 8 = 10

10 % 8 = 2
ans = "2" + "3"
10 // 8 = 1

1 % 8 = 1
ans = "1" + "2" + "3"
1 // 8 = 0

ans = 123
"""
