n,x=map(int,input().split())
x-=1
alp='abcdefghijklmnopqrstuvwxyz'
ans=alp[x//n]
print(ans.upper())
