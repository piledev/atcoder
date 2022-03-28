"""
【方針】 問題を簡単なものにすり替えていく。

1.「総和がNになる数列の数を数える」

公差1の等差数列であることは決まっているので

2.「総和がNになる初項aと末項bの組み合わせを数える」

(a+b)*(b-a+1) == 2n が成立する。
ここで 2n == x*y(x,yは整数) として以下のように割り当ててみる。
① : x==a+b
② : y==b-a+1

3. 「式① ② を満たす整数a, bの組み合わせを数える」

これらの形を変えると以下のようになる。
③ : a==(x-y+1)/2
④ : b==(x+y-1)/2
a, b がともに整数であるためには
(x-y+1)も、(x+y-1)も、偶数でなければならない。
x, y の偶奇が一致するとこれは成立しない。

また① ② より、 x,yが一意に決まればa,bも一意に決まりそう。
なので以下が成立する。

4.「2N==x*yと分解したとき、x, yの偶奇が異なる組み合わせを数える」
"""
n=int(input())
ans=0
x=1
while x**2<=2*n:
    if 2*n%x==0 and (2*n//x)%2 != x%2:
        ans+=2
    x+=1
print(ans)

