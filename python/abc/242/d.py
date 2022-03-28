s=str(input())
q=int(input())
abc=["A","B","C"]
def rec(t,k):
    if t==0:
        if s[k-1]==abc[0]:
            return 0
        elif s[k-1]==abc[1]:
            return 1
        else:
            return 2
    if k%2==1:
        return (rec(t-1,(k+1)//2)+1)%3
    return (rec(t-1,k//2)+2)%3

for i in range(q):
    t,k=map(int,input().split())
    ans=rec(t,k)
    print(abc[ans])


