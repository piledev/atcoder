x,k,d = map(int,input().split())
x=abs(x)

# k回移動しても0に到達しない場合
if 0<x-k*d:
    print(abs(x-k*d))

# 到達する場合
else:
    # x//d: 0 を飛び越える直前までの移動回数
    # move_count: 0を飛び越える直前の残り移動回数
    move_count = k-x//d
    # 飛び越える直前の位置の絶対値
    jump_before = abs(x-d*(x//d))
    # 飛び越えた直後の位置の絶対値
    jump_after = abs(jump_before-d)
    # 残り回数が偶数なら直前の位置に戻ってくる
    if move_count%2==0:
        print(jump_before)
    # 残り回数が奇数なら直後の位置
    else:
        print(jump_after)

