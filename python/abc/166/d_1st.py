x=int(input())
# ↓でx=200とした時点で大きくなりすぎることを突き止める
# print(x**5-(x-1)**5)

for a in range(-200,200):
    for b in range(-200,200):
        ans= a**5-b**5
        if ans==x:
            print(a,b)
            exit()

