n=int(input())

t=''
for i in range(1,n+1):
    if i==1:
        t='1'
    else: 
        t=t+' '+str(i)+' '+t

print(t)
    
