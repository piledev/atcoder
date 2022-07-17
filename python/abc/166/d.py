x=int(input())

for a in range(-1000,1000):
    for b in range(-1000,1000):
        if x==a**5-b**5:
            print(a,b)
            exit()
