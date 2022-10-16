h1,w1=map(int,input().split())
a=[]
for i in range(h1):
    r=list(map(int,input().split()))
    a.append(r)
h2,w2=map(int,input().split())
b=[]
for i in range(h2):
    r=list(map(int,input().split()))
    b.append(r)
y='Yes'
n='No'

ans=n

def match(a,b):
    ans=y
    print('a:',a)
    if len(a)!=len(b) or len(a[0])!=len(b[0]):
        return False 
        
    for i in range(len(b)):
        for j in range(len(b[i])):
            if a[i][j] != b[i][j]:
                return False 
    return True 

for i in range(1<<h1):
    for j in range(1<<w2):
        a_reduced=list()
        for hshift in range(h1):
            row=list()
            for wshift in range(w1):
                if i>>hshift & 1 == 1 and j>>wshift & 1 == 1:
                    row.append(a[hshift][wshift])
            if len(row)>0:
                a_reduced.append(row)
        if match(a_reduced,b):
            print(y)
            exit()

print(n)



                    





