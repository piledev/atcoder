n=int(input())
mods=[0]*3

t=n
while t>=10:
    mods[t%10%3]+=1
    t//=10
mods[t%10%3]+=1

mod=n%3
if mod == 0:
    print(0)
elif mod == 1:
    if mods[1]>0 and sum(mods)>1:
        print(1)
    elif mods[2]>1 and sum(mods)>2:
        print(2)
    else:
        print(-1)
elif mod == 2:
    if mods[2]>0 and sum(mods)>1:
        print(1)
    elif mods[1]>1 and sum(mods)>2:
        print(2)
    else:
        print(-1)




