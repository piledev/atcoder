# coding: utf-8
import random

# global変数はアッパーケースで定義するのが慣習
GLOBAL_VALUE = "Global"

# 割り算と助産
print("割り算:",10/3)
print("除算  :",10//3)

# 文字列も繰り返せる
print("hello"*3)

# 文字列の埋め込み
name = "yokoyamada"
score = 100 
print("name: %s, score: %d" %(name,score))
# 右揃えや左揃えもある
print("name: %12s, score: %5d" %(name,score))
print("name: %-12s, score: %-5d" %(name,score))

print("")
print("=====================================================")

# 文字列の埋め込み２
print("name: {0}, score: {1}".format(name,score))
# 桁数の指定
print("name: {0:12s}, score: {1:5d}".format(name,score))
# >右揃えや<左揃えもある
print("name: {0:>12s}, score: {1:<5d}".format(name,score))

print("")
print("=====================================================")

# if文
if score > 80:
    print("Great!")
elif score > 60:
    print("Good")
else:
    print("So so...")

# 条件演算子
print("Great!" if score > 80 else "So so...")

print("")
print("=====================================================")

# While 文
i = 0
while i < 10:
    if i == 5:
        break
    print("%d:%s"%(i,name))
    i+=1
else:
    print("finish!")

print("")
print("=====================================================")

# for 文とloop終了後の処理(else)
for i in range(0,5):
    if i%2==0  :
        continue
    print(i)
else:
    print("finish!")

print("")
print("=====================================================")
# コレクション
# list
list1 = [10,20,30,40]
list2 = [10,20,40]
list3 = "Hi!!" # 文字列もリスト
print("list : {0}".format(list1))

# tuple
tuple1 = 10,20,30,40
print("tuple: {0}".format(tuple1))

# dicti
dict1 = {"dog":"犬", "cat":"猫"}
print("dict: {0}".format(dict1))
key1 = "dog"
print("{0} は {1} です。".format(key1, dict1[key1]))

# 大小比較もできる(list, tuple)(setの比較基準はなんか違うっぽい)
print("{0} < {1} =".format(list1,list2),list1 < list2)

# 添字指定で編集できる(list)
list1[2] = 31

# 添字指定で編集できない(tuple, dict)
# tuple1[2] = 31

# tuple の更新をする場合は変数そのものを上書きする必要がある
tuple1 = 1,2

# for ~ in(list, tuple, dict, set)
for v in list1:
    print("value: {0}".format(v))

# enumerate で添字も取れる(list, tuple, set)
for i, v in enumerate(list1):
    print("list[{0}]: {1}".format(i,v))

# dict の場合は items()を使えばkeyもvalueも取れる
for k, v in dict1.items():
    print("dict[{0}]: {1}".format(k,v))

# in 演算子（要素を含むかどうか）(list, tuple, dict, set)
print(30 in list1)
print("!" in list3) # 文字列は複数指定できる
print("cat" in dict1) # dict はvalueではなくkeyを指定する

# 要素の追加(list)
print(list1)
list1.append(100)
print(list1,"< appended.")

# list と tuple の双方変換
print(tuple(list1))
print(list(tuple1))

# slice(goのsliceは型だけど、pythonでは動詞)
print(list1[:])
print(list1[:2])
print(list1[3:])
# - は「末尾から数えて」という意味になる
print(list1[-3:-1])

# アンパック 
a, b, c, d, e = list1
# 変数の数が要素数と異なる場合はエラー
# a, b, c  = list1
# dict のときはkeyを取得
a, b = dict1
print(a, b)
# アンパックの活用例
# 複数の戻り値を返したいとき
def sum_and_average(num_list):
    total = 0
    for v in num_list:
        total += v
    average = total / len(num_list)
    # タプルで返す
    return (total,average)
# 受け取りつつアンパック
sum1, ave1 = sum_and_average(list1)
print(sum1, ave1)

print("")
print("=====================================================")
# set
"""
set は重複を許さなくて順序もないデータ型
{}で囲むと出来上がる
"""
set1 = {5,4,8,5}
# このようにしてリストから重複を除くという使い方もできる(順序はなくなるけど)
set2 = set([5,4,9,5])
print(set1)
print(set2)

# 要素の追加
set1.add(9)
print(set1)
# 要素の削除
set1.remove(8)
print(set1)

# 比較結果は要素が完全一致するかしないか(ぽいけどどうだろう？)
print(set1 == set2)

setx = {1,3,5,8}
sety = {3,5,4,9}

print("setx: {0}".format(setx))
print("sety: {0}".format(sety))

# 和集合
print("和集合(|)  : {0}".format(setx | sety))
# 積集合
print("積集合(&)  : {0}".format(setx & sety))
# 差集合
print("差集合(x-y): {0}".format(setx - sety))
print("差集合(y-x): {0}".format(sety - setx))

# set も enumerate で添字も取れるが順不同
for i,v in enumerate(setx):
    print(i,v)
   
print("")
print("=====================================================")
# イテレータ
"""
イテレータはどこまでの要素を取得したのかを覚えていてくれているもの
next(iter) で次の要素を取得する
"""
score = [10,20,40,80]
it = iter(score)
print(next(it))
print("sorry to disturb you, but...")
print(next(it))

print("")
print("=====================================================")
# ジェネレータ
"""
イテレータを作る関数をジェネレータと呼ぶ。
yield を使うことでジェネレータを実装できる。
listでは扱えないような無限の要素を持つイテレータを作ったりできる
"""

# ジェネレータ（イテレータを作る関数）
def get_infinite():
    i = 0
    while True:
        # yield は関数を一時的に実行停止し、一旦の戻り値を返す
        yield i*2 
        i += 1

generator = get_infinite()
print(next(generator))
print(next(generator))
print(next(generator))

def get_three():
    yield "one"
    yield "two"
    yield "three"
generator = get_three()
print(next(generator))
print(next(generator))
print(next(generator))
# 4回目にはyield はもう終わっているのでエラーになる
# print(next(generator))

# generator は統合できる
def get_three_and_more(g1,g2):
    yield from g1
    yield from g2

generator = get_three_and_more(get_three(),get_infinite())
print(next(generator))
print(next(generator))
print(next(generator))

print("")
print("=====================================================")
# map
"""
- map() はイテレータの要素を加工するのに使う
    - イテレータの各要素に対して同じ処理を施して、
      同じ長さの要素にすべての結果を格納して返してくれる関数
    - 戻り値はジェネレータなので、cast して使うのが吉？
      ans = list(map(~))
"""
def triple(n):
    return n*3
print(list(map(triple,[1,2,3])))

print("")
print("=====================================================")
# lambda(無名関数)
lambdafunc = lambda x:x*2
print(lambdafunc(2))
print(lambdafunc(4))

print("")
print("=====================================================")
# lambda + map
list2 = list(map(lambda x:x*5,list1))
print(list1)
print(list2)

print("")
print("=====================================================")
# lambda + sorted()
tuplelist1 = [(7, 2), (3, 4), (5, 5), (10, 3)]
# tuple の2番目の要素でソートするという使い方
tuplelist2 = sorted(tuplelist1,key=lambda x:x[1])
print(tuplelist1)
print(tuplelist2)

print("")
print("=====================================================")
# lambda + filter()
class Album(object):
    def __init__(self, title, artist):
        self.title = title
        self.artist = artist
a1 = Album("A Hard Day's Night", "The Beatles")
a2 = Album("The Rolling Stones", "The Rolling Stones")
a3 = Album("Abbey Road", "The Beatles")
albums = [a1, a2, a3]

albums_beatles = filter(lambda album: album.artist == "The Beatles", albums)
for a in albums_beatles:
    print(a.artist, "-", a.title)

# lambda ではないけどこんなのもいける
def is_even(n):
    return n % 2 == 0
print(list(filter(is_even,range(20))))

print("")
print("=====================================================")
# 内包表記
"""
リストやジェネレータを生成したり加工する際に書くための記法
↓のコードの意味はrange(10)で0から9のリストができるので、
そのリストから1つずつ要素を取り出してiに入れて、
それを二乗して取り出すという意味になる。
ｍap() でも同様の処理ができるが、
リスト内包表記のほうがコードが簡潔・明解で好ましいとされている
"""
print([i**2 for i in range(10)])
# set型でやることも可能(順不同・重複削除になる)
print({i**2 for i in range(10)})

"""
これをfor文で書くとこう
"""
squares = []
for i in range(10):
    squares.append(i**2)
print(squares)

"""
こういうのもある
filter() でも同様の処理ができるが、
リスト内包表記のほうがコードが簡潔・明解で好ましいとされている
"""
adds = [i for i in range(10) if i%2 == 1]
print(adds)
    
exit()
print("")
print("=====================================================")

# 関数(引数にはデフォルト値を設定している)
def say_hi(name="default Name",age=99):
    print("Hi! {0} age({1})".format(name,age))

# 関数呼び出しは必ず関数の定義後
say_hi(name,44)
# デフォルト引数を使う
say_hi()
# 引数の順番は自由
say_hi(age=80,name="freedom")

print("")
print("=====================================================")

# 関数の中身の実装は保留することができる。Noneを返す。
def not_implemented_yet():
    pass

ret = not_implemented_yet()
print(ret)


print("")
print("=====================================================")

msg = "Global!" # グローバル変数

def local_func():
    # msg はローカル変数
    msg = "change in local"
    print(msg)

local_func()
print(msg) # グローバル変数は関数からは変更できない

def change_global_variable():
    # グローバル変数のmsgを明示的に宣言
    global msg
    msg = "change in local"
    print(msg)

change_global_variable()
print(msg) 

print("")
print("=====================================================")

# class
# クラス名はアッパーキャメルケースが慣習
# class のプロパティは外から追加できる（気持ち悪いけど）
# インスタンス固有のプロパティにはなってしまうけど
class User1:
    pass

tom = User1()
tom.name = "Tom"
tom.score = 20

bob = User1()
bob.name = "Bob"
bob.level = 100

print(tom.name, tom.score)
print(bob.name, bob.level)

# class (コンストラクタを使った気持ちいい使い方)
class User2:
    # self はこのクラスから作られるインスタンスを指す
    # self は予約語でも何でもなくてただの変数名。違う名前にしても大丈夫。
    # __init__は予約後っぽい。第一引数がインスタンスになる。
    def __init__(self,name):
        self.name = name

tom = User2("Tom")
bob = User2(12)

print(tom.name)
print(bob.name) # 当然文字列じゃなくても入る

print("")
print("=====================================================")

# クラス変数とインスタンス変数
class User3:
    count = 0
    # コンストラクタ
    def __init__(self,name):
        # クラス変数
        User3.count += 1
        # インスタンス変数
        self.name = name

print(User3.count)
tom = User3("Tom")
bob = User3("Bob")
print(User3.count)

# クラスのプロパティなのでインスタンスからも呼び出せる
print("tom is count {0}".format(tom.count))

print("")
print("=====================================================")

# メソッド
# クラスに定義した関数をメソッドという

class User4:
    count = 0
    def __init__(self,name):
        User4.count +=1
        self.name = name

    # インスタンスメソッド
    def say_hi(self):
        print("hi! {0}.".format(self.name))

    # クラスメソッド
    """
    クラス変数(cls)をそのまま使えるメソッド。
    インスタンスを生成しないでクラス内のメソッドを呼び出せる。
    デコレータ指定されたクラスメソッドは、呼び出される際に、暗黙の引数としてクラス自身が渡される。
    必ず渡されるので、その受け皿をcls（変数名は何でもいいんだけど）として定義しておく必要がある。
    """
    @classmethod # ←デコレータ
    def show_info(cls):
        print("{0}:instances".format(cls.count))

    # だめな例
    # @classmethod # ←デコレータ
    # def show_info():
        # print("{0}:instances".format(3))
    # ↓ 
    # TypeError: User4.show_info() takes 0 positional arguments but 1 was given
 
        
User4.show_info()
tom = User4("Tom")
User4.show_info()


print("")
print("=====================================================")

# Python には private や public 的なアクセス権を明示的に宣言する仕組みが無いので、
# 暗黙のルールとして変数名の前に__をつけることでアクセス制限できる。
class User5:
    count = 0
    def __init__(self,name):
        User5.count += 1
        ## クラス外からアクセスしないように__をつける
        self.__name = name
    def say_hi(self):
        print("hi! {0}".format(self.__name))

tom = User5("Tom")
# ↓これはエラーになる
# print(tom.__name)
# でも裏技で見える
print(tom._User5__name)

print("")
print("=====================================================")
# クラスの継承
# 親クラス
class User6:
    def __init__(self,name):
        self.name = name
    def say_hi(self):
        print("[親クラスのメソッド] Hi! {0}".format(self.name))

# 子クラス
class ChildUser6(User6):
    def __init__(self,name,age):
        """
        親クラスのコンストラクタで name を初期化する処理がすでにあるので、
        それを子クラスの中でも再利用するために super() を使う。
        これを呼び出すことで`self.name`が使える。
        つまりここは、親クラスのself.nameと同じ場所を参照している。
        """
        super().__init__(name)
        self.age = age
        print("子クラスでsuper()を使った親クラスのプロパティ呼び出されるnameは {0}".format(self.name))
        self.name = "ChildUser6"
        print("子クラスで親クラスのプロパティを書き換えたあとに呼び出されるnameは {0}".format(self.name))

    def say_hello(self):
        print("Hello!! {0} age:{1}".format(self.name,self.age))

    def say_hi(self):
        print("[子クラスでオーバーライドされたメソッド] Hi! {0}".format(self.name))




print("継承前の世界")
bob = User6("Bob")
bob.say_hi()

print("継承後の世界")
bob = ChildUser6("bob", 22)
bob.say_hi()

print(bob.name)
bob.say_hello()
bob.say_hi()



print("")
print("=====================================================")
# 多重継承
class A:
    def say_a(self):
        print("A")
    def say_hi(self):
        print("Hi from A")
class B:
    def say_b(self):
        print("B")
    def say_hi(self):
        print("Hi from B")
class C(B,A):
    pass

c = C()
c.say_a()
c.say_b()
# 同じメソッド名を保つ場合はC(B,A)の()の中で先に宣言されているBが優先される
c.say_hi()

print("")
print("=====================================================")
# モジュールのインポート
# モジュールはファイル名で指定する
# 関数やクラスを指定してインポートできる
# ,区切りで複数クラスの指定もできる
# インポートしてないものは見えない
from ex_module import ExChildUser, ExUser
tom = ExChildUser("tom",23)

# モジュールそのものもインポートできる
import ex_module
# その場合はモジュール名.~ で指定して使う
tom = ex_module.ExUser("Tom")

print("")
print("=====================================================")
# パッケージモジュール
"""
パッケージはフォルダがモジュール化したものと考えて良い。
パッケージフォルダの中には、__init__.py が必要
"""
import mypackage.package_module
tom = mypackage.package_module.PxChildUser("Tom", 23)
# 別名つけることもできる
import mypackage.package_module as mpack
bob = mpack.PxChildUser("Bob", 23)
# 特定の関数のみ読み込むこともできる
from mypackage.package_module import PxChildUser

print("")
print("=====================================================")
# 例外処理

# 基本構文
"""
# 用意されている例外クラスを使う場合

def Hoge()
    try:
        # 例外が起こりそうな処理
    except 例外が起きた際に呼び出されるクラス名:
    else:
        # 例外が起きなかった場合の処理
    finally:
        # 例外関係なく実行してほしい処理
"""
"""
# 独自で例外クラスを作る場合

class MyException(Exception):
    pass
def hoge():
    try: 
        # 任意の条件で例外を投げる
        raise MyException("My Exception!!!")
    # 独自例外発生時の処理
    except MyException as e:
        # e には自分で作った独自処理のエラーメッセージが入っている
        print(e)

"""
class MyException(Exception):
    pass

def div(a, b):
    try:
        if (b < 0):
            raise MyException("minus!")
        print(a/b)
    except ZeroDivisionError:
        print("not by zero...")
    except MyException as e:
        print(e)
    else:
        print("no exception!")
    finally:
        print("end.")
div(10,3)
div(10,0)
div(10,-1)

print("")
print("=====================================================")

exit()
print("")
print("=====================================================")

print("")
print("=====================================================")

print("")
print("=====================================================")

print("")
print("=====================================================")

print("")
print("=====================================================")

