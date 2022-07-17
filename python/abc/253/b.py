h,w=map(int,input().split())
start_h,start_w=-1,-1
goal_h,goal_w=-1,-1
for i in range(h):
    s=list(str(input()))
    for j in range(w):
        if s[j]=='o':
            if start_h==-1:
                start_h=i
                start_w=j
            else :
                goal_h=i
                goal_w=j

ans=abs(goal_h-start_h)+abs(goal_w-start_w)
print(ans)


