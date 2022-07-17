n, k=map(int,input().split())
r,s,p=map(int,input().split())
t=input()

history=["x" for i in range(n)]
stronger={"r":"p","s":"r","p":"s"}

def janken(score, my):
    if my=="r":
        score+=r
    elif my=="s":
        score+=s
    elif my=="p":
        score+=p
    return score, my

ans=0
for i in range(n):
    hand=t[i]
    if i < k:
        ans, history[i] = janken(ans, stronger[hand])
    else:
        if history[i-k] != stronger[hand]:
            ans, history[i] = janken(ans, stronger[hand])

print(ans)

