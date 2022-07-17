n,k=map(int,input().split())
s=list()
for i in range(n):
    s.append(list(input()))

ans=0
for i in range(1<<n):
    counter=dict()
    for shift in range(n):
        if i>>shift&1==1:
            for c in s[shift]:
                if c not in counter:
                    counter.update({c:0})
                counter[c]+=1
    count=0
    # print(counter)
    for key,val in counter.items():
        # print(key,val)
        if val==k:
            count+=1
    ans=max(ans,count)

print(ans)



                


