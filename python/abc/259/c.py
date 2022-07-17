s=input()
t=input()

scnts=list()
tcnts=list()

scnts.append([s[0],1])
for i in range (1,len(s)):
    if s[i]==scnts[-1][0]:
        scnts[-1][1]+=1
    else:
        scnts.append([s[i],1])

tcnts.append([t[0],1])
for i in range (1,len(t)):
    if t[i]==tcnts[-1][0]:
        tcnts[-1][1]+=1
    else:
        tcnts.append([t[i],1])

if len(scnts) != len(tcnts):
    print('No')
    exit()

for i in range(len(scnts)):
    if scnts[i][0] != tcnts[i][0]:
        print('No')
        exit()
    else:
        if scnts[i][1] > tcnts[i][1]:
            print('No')
            exit()
        elif scnts[i][1] == tcnts[i][1]:
            continue
        elif scnts[i][1]>=2:
            continue
        else:
            print('No')
            exit()

print('Yes')
