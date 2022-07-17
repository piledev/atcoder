n=int(input())

insu_cnt=list()
for a in range(1,7):
    b=1
    a_insu_cnt=0
    while b**2 < a:
        if a%b==0:
            print("a,b:",a,b)
            a_insu_cnt+=2
        b+=1
    a_insu_cnt+=1
    insu_cnt.append((a**2,a_insu_cnt))
    print((a**2,a_insu_cnt))




