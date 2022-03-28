'''
1個の場合条件を満たすか？2個の場合は？...と順に確かめる。
x個の場合には以下を満たせば良い。
a*x <= 1000*w <= b*x
(↑この式が思いつかなかったし、思いつける気がしない。)

w の最大値は1000kg = 10**6g
k の最小値は1g なので、最大でも10**6 回の計算となり、間に合う。
'''

a,b,w = map(int,input().split())
w*=1000
min_ans = 10**15
max_ans = -10**15

for x in range(1,w//a+1):
    if a*x <= w <= b*x:
        min_ans = min(min_ans,x)
        max_ans = max(max_ans,x)

if min_ans == 10**15:
    print("UNSATISFIABLE")
else:
    print(min_ans,max_ans)

