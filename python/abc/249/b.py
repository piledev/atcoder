s=list(input())
upper=False
lower=False

sets=set(s)
if len(s)==len(sets):
    for c in sets:
        if c.isupper():
            upper=True
        else:
            lower=True
    if upper and lower:
        print('Yes')
        exit()

print('No')

