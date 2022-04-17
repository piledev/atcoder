import itertools

# 順列
"""
(0, 1, 2),(0, 2, 1),(1, 0, 2),(1, 2, 0),(2, 0, 1),(2, 1, 0)
"""
for seq in itertools.permutations(range(3)):
    print("{0},".format(seq),end='')
print()

# 重複なしの組み合わせ
"""
(0, 1),(0, 2),(0, 3),(1, 2),(1, 3),(2, 3)
"""
for seq in itertools.combinations(range(4), 2):
    print("{0},".format(seq),end='')
print()

# 重複ありの組み合わせ
"""
(0, 0),(0, 1),(0, 2),(0, 3),(1, 1),(1, 2),(1, 3),(2, 2),(2, 3),(3, 3)
"""
for seq in itertools.combinations_with_replacement(range(4), 2):
    print("{0},".format(seq),end='') 
print()

# 直積
"""
(0, 1),(0, 2),(1, 1),(1, 2),(2, 1),(2, 2),(3, 1),(3, 2)
"""
for seq in itertools.product(range(4),range(1,3)):
    print("{0},".format(seq),end='') 
print()

