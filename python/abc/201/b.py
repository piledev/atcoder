n = int(input())
a = []
for i in range(n):
    # name, height = input().split()
    # a.append((name,int(height)))
    name, height = map(str,input().split())
    height = int(height)
    a.append((height, name))

# a_s = sorted(a, key=lambda x:x[1],reverse=True)
a.sort(reverse=True)
print(a[1][1])
