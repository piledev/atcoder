# 22:40
n, k = map(int,input().split())
friends = []
for i in range(n):
    a, b = map(int,input().split())
    friends.append((a,b))

friends.sort()
nextfriend = 0 
current = k 
while friends[nextfriend][0] <= current:
    current += friends[nextfriend][1]
    nextfriend += 1
    if nextfriend == n:
        break
print(current)


