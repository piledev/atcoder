UNSATISFIABLE = "UNSATISFIABLE"

def calc_min(a,b,w):
    bmod = w%b
    bcnt = w//b
    if bmod == 0:
        return bcnt 

    cnt = bcnt
    mod = bmod
    bquot = mod//b
    aquot = mod//a
    while mod <= w: 
        if bquot != aquot:
            return cnt+aquot
        cnt-=1
        mod+=b
        bquot=mod//b
        aquot=mod//a
    return 0

def calc_max(a,b,w):
    amod = w%a
    acnt = w//a
    if amod == 0:
        return acnt 

    cnt = acnt
    mod = amod
    aquot = mod//a
    bquot = mod//b
    while mod <= w: 
        if bquot != aquot:
            return cnt+aquot
        cnt-=1
        mod+=a
        bquot=mod//b
        aquot=mod//a
    return 0

a,b,w = map(int,input().split())
w*=1000
mincnt=calc_min(a, b, w)
if mincnt == 0:
    print(UNSATISFIABLE)
    exit()
maxcnt=calc_max(a, b, w)
print(mincnt,maxcnt)

