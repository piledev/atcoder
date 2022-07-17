n,q=map(int,input().split())
balls=[i+1 for i in range(n)]
indexes={i+1:i for i in range(n)}

for i in range(q):
    x=int(input())
    swap_index=0
    x_index=indexes[x]
    if x_index==n-1:
        swap_index=n-2
    else:
        swap_index=indexes[x]+1
    swap_ball=balls[swap_index]

    balls[x_index], balls[swap_index]=swap_ball, x
    indexes[x], indexes[swap_ball]=swap_index,x_index

print(' '.join(map(lambda x:str(x), balls)))


    

