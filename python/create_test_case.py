path = "./test.txt"
s="1"

for i in range(10**4-1):
    s+=" 1"

with open(path,mode="w") as f:
    f.write(s)
