def expected_value(x):
    return((1+x)/2)

n, k = map(int,input().split())
p = list(map(int,input().split()))
expected_sum = int()
expected_values = list()

for i in range(k):
    temp = expected_value(p[i])
    expected_values.append(temp)
    expected_sum += temp

max_expected = expected_sum
for i in range(k,n):
    temp = expected_value(p[i])
    expected_values.append(temp)
    expected_sum += temp
    expected_sum -= expected_values[i-k]
    max_expected = max(max_expected,expected_sum)

print(max_expected)

