import math
n,a,b=map(int,input().split())
gcd=math.gcd(a,b)
lcm=a*b//gcd

n_sum=(n+1)*n//2
n_div_a=n//a
n_div_b=n//b
n_div_lcm=n//lcm
a_sum=a*((n_div_a+1)*n_div_a//2)
b_sum=b*((n_div_b+1)*n_div_b//2)
lcm_sum=lcm*((n_div_lcm+1)*n_div_lcm//2)

ans=n_sum-(a_sum+b_sum-lcm_sum)
print(ans)
