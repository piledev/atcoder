# RULES
onsight : AC on first attempt without any previous information. 
flash:    AC on first attempt with some pre-existing knowledge.
redpoint: AC after having practiced.

# ABC
## 205
### Result
```
a: AC onsight
b: AC onsight
c: AC onsight
d: TLE
e: -
f: - 
```
### Technique

## 196
### Result
```
a: AC onsight 
b: AC onsight
c: AC redpoint
d: TLE
e: -
f: - 
```
### Technique
- DFS


## 206
### Result
```
a: AC onsight
b: AC onsight
c: AC redpoint
d: AC redpoint (but time out...)
e: -
f: - 
```
### Technique
- Union-Find
- DFS
- BFS

## 194 
### Result
```
a: AC onsight
b: AC onsight
c: AC redpoint
d: -  I can understand editorial.
e: -  I can understand editorial.
f: -  I can't understand editorial.
```
### Technique
- 有効なのが来るまでカードを引く期待値は、有効なカードを引く確率の逆数になる。
- 確率 p で成功する試行を成功するまで行うときの試行回数（最後の成功した回も含む）の期待値は 1/p である。
- DP法

## 193 
### Result
```
a: AC onsight
b: AC onsight
c: AC redpoint
d: -  I can't understand editorial.  
e:   
f:   
```
### Technique
-

## 207
### Result
```
a: AC onsight
b: AC redpoint (but time out...)
c: AC redpoint
d: -  I can't understand editorial.  
e:
f: 
```
### Technique
- int で math.Ceil(x/y) を実装する方法
-- 分子に「分母-1」を加えてやることで、math.Ceil と同じ効果を得られる。
-- ex. ans := (x+y-1)/y

## 208
### Result
```
a: AC onsight
b: AC onsight
c: AC onsight
d: -  I can't understand editorial.  
e:
f: 
```
### Technique


## 209
### Result
```
a: AC redpoint
b: AC onsight
c: AC onsight
d: -  I can understand editorial.  
e:
f: 
```
### Technique
- BFS
- queueの使い方
```go
// 1. allocate のオーバーヘッドを避けるため、必要な領域を確保しよう。
queue := make([]int,N)

// 2. 初期値を追加しよう.
quere = append(quere, 0)

// 3. キューの追加を行いながら、キューがなくなるまで処理しよう
for len(queue) != 0 {
  // 4.キューの先頭の値を取り出し、詰めよう。
  t:= queue[0]
  queue = queue[1:]
  ...
  for ... {
    queue = append(queue, x)
  }

}
```
- 色の反転(2色の場合)
```go
// color: 0 or 1 
// 1. bit演算子(xor)を使用する
color ^= 1
// 2. 普通の引き算
color = 1-color
```

# Educational DP Contest 
https://atcoder.jp/contests/dp
## A: Frog 1
### Result
```
2021/06/27: AC on "morau" & "kubaru". GU on "recursive".
```
### Technique

## B: Frog 2 
### Result
```
2021/06/27: AC on "morau" & "kubaru" & "recursive" 
```

## C: Vacation 
### Result
```
2021/06/27: AC on "morau" & "kubaru" & "recursive" 
```
### memo 
- 時間さえかければ再帰も書ける。
- 再帰関数の中に再帰関数の呼び出しを書き忘れる。
- 多次元配列の各次元が何を意味するのかわからなくなる。
- 「index 0 のときは...」とか考えたくないので、dp のindex 0 は初期値 0 を設定する場所にしたい。だから要素数は原則 N+1 にしよう。

## D: Knapsack 1 
### Result
```
2021/06/27: 
```

## E: Knapsack 2 
### Result
```
2021/06/27: 
```

## F: LCS 
### Result
```
2021/06/27: 
```

