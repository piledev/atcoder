n=int(input())
a=list(map(int,input().split()))
MAX=200000

# list に対して in すると O(n) かかるらしいので、
# O(n) で in できる dict を使用することにした。
# しかし、i と j の二重ループを解消することができず、
# 結局のところTLEする結果となった。
# d=dict()
# for i in a:
#     if i not in d:
#         d.update({i:0})
#     d[i]+=1
# ans=0
# for i,iv in d.items():
#     for j,jv in d.items():
#         if i*j in d:
#             ans+=iv*jv*d[i*j]
# print(ans)

# 以下が公式解説通りの解法。
# ポイント1)
# a[i]=a[j]*a[k]
# a[i]<=200000
# であることを利用し、a[j]を探す範囲を絞った。
# ポイント2)
# index に a[x] の値を取る list c を作成した。
# c の値は x の出現数。
# i,j,k の組み合わせの数 =  a[i] の出現数 * a[j]の出現数 * a[k]の出現数
# となることと合わせて、
# ans+=c[i]*c[j]*c[i*j] と単純化することができた。
# (i,j,kのうちどれか１つでも出現していなかったら0になる。)

c=[0 for i in range(MAX+1)]
for i in a:
    c[i]+=1

ans=0
for i in range(1,MAX+1):
    j=1
    while i*j<=MAX:
        ans+=c[i]*c[j]*c[i*j]
        j+=1

print(ans)
