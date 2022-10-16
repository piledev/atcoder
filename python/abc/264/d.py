dic={
'a':0,
't':1,
'c':2,
'o':3,
'd':4,
'e':5,
'r':6
}
s=list(input())
si=list()
for i in range(7):
    si.append(dic[s[i]])

def BubbleSort(num):
    count=0
    for i in range(len(num)):
        for j in range(len(num)-1, i, -1):
            if num[j] < num[j-1]:
                num[j], num[j-1] = num[j-1], num[j]
                count+=1

    return count 

ans=BubbleSort(si)
print(ans)
