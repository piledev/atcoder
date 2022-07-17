a,b,c,d,e,f,x=map(int,input().split())

def abcx(a,b,c,x):
    ret=x//(a+c)*a*b
    if x%(a+c)<a:
        ret+=x%(a+c)*b
    else:
        ret+=a*b
    return ret 

takahashi=abcx(a,b,c,x)
aoki=abcx(d,e,f,x)

ans=''
if takahashi>aoki:
    ans='Takahashi'
elif aoki>takahashi:
    ans='Aoki'
else:
    ans='Draw'

print(ans)

