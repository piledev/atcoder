s=str(input())
q=int(input())

# g(): s を0番目として数えて add 番目の文字を返す
# s  : 起点となる文字
# add: 番目
def g(s,add):
    return chr((ord(s)-ord('A')+add)%3+ord('A'))

# f(): t回変換後のk文字目を取得する
def f(t,k):
    # 0 文字目は A,B,C,... のように循環している
    if k==0:
        return g(s[0],t)
    # 変換前の文字
    if t==0:
        return s[k]

    # 再帰的に今回の変換前の文字を取得
    bef = f(t-1,k//2)

    # kが偶数のときは次の文字（AならB,BならC）
    # kが奇数のときは次の次の文字（AならC,BならA）
    return g(bef,k%2+1)

for i in range(q):
    t,k=map(int,input().split())
    print(f(t,k-1))

