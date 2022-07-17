n,a,b=map(int,input().split())

white='.'
black='#'
cells=[[white for i in range(b*n) ] for j in range(a*n)]

for i in range(a*n):
    for j in range(b*n):
        if (i//a)%2==0:
            if (j//b)%2==1:
                cells[i][j]=black
        else:
            if (j//b)%2==0:
                cells[i][j]=black

for v in cells:
    print(''.join(v))
